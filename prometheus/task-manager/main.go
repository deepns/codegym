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

// statusRecorder wraps http.ResponseWriter to capture response status code
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

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

	// Error counter by HTTP status code
	httpErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors by status code",
		},
		[]string{"method", "endpoint", "status"},
	)
)

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	taskID++
	task := Task{ID: taskID, Name: "Task-" + fmt.Sprint(taskID)}
	tasks[taskID] = task
	activeTasks.Inc()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
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
}

func main() {
	// Register Prometheus metrics
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(activeTasks)
	prometheus.MustRegister(httpErrors)

	// Define metrics middleware
	metricsMiddleware := func(endpoint string, next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// default status 200 in case handler doesn't call WriteHeader explicitly
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

			next.ServeHTTP(rec, r)

			method := r.Method
			// record request count and duration
			httpRequests.WithLabelValues(method, endpoint).Inc()
			requestDuration.WithLabelValues(method, endpoint).Observe(time.Since(start).Seconds())

			// record errors (status >= 400)
			if rec.status >= 400 {
				httpErrors.WithLabelValues(method, endpoint, fmt.Sprint(rec.status)).Inc()
			}
		})
	}

	// Define routes
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/tasks", metricsMiddleware("/tasks", http.HandlerFunc(listTasksHandler)))
	http.Handle("/tasks/create", metricsMiddleware("/tasks", http.HandlerFunc(createTaskHandler)))
	http.Handle("/tasks/delete", metricsMiddleware("/tasks", http.HandlerFunc(deleteTaskHandler)))

	// Start server
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
