package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	tasks     = make(map[int]Task)
	taskMutex sync.Mutex
	taskID    = 0

	// Prometheus Metrics
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for requests",
			Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
		},
		[]string{"method", "endpoint"},
	)

	activeTasks = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_tasks",
			Help: "Current number of active tasks",
		},
	)
)

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	httpRequests.WithLabelValues("POST", "/tasks").Inc()

	taskMutex.Lock()
	defer taskMutex.Unlock()

	taskID++
	task := Task{ID: taskID, Name: "Task-" + fmt.Sprint(taskID)}
	tasks[taskID] = task
	activeTasks.Inc()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

	requestDuration.WithLabelValues("POST", "/tasks").Observe(time.Since(start).Seconds())
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	httpRequests.WithLabelValues("GET", "/tasks").Inc()

	taskMutex.Lock()
	defer taskMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

	requestDuration.WithLabelValues("GET", "/tasks").Observe(time.Since(start).Seconds())
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	httpRequests.WithLabelValues("DELETE", "/tasks").Inc()

	taskMutex.Lock()
	defer taskMutex.Unlock()

	if len(tasks) > 0 {
		for id := range tasks {
			delete(tasks, id)
			activeTasks.Dec()
			break
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task deleted"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No tasks found"))
	}

	requestDuration.WithLabelValues("DELETE", "/tasks").Observe(time.Since(start).Seconds())
}

func main() {
	// Register Prometheus metrics
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(activeTasks)

	// Define routes
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/tasks", listTasksHandler)
	http.HandleFunc("/tasks/create", createTaskHandler)
	http.HandleFunc("/tasks/delete", deleteTaskHandler)

	// Start server
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
