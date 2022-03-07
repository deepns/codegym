// Revisiting some concepts learnt earlier
// json decode
// http Get
// go routines
// timers
// select
// channels
//	send
//	receive
// defer

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// demoGet makes a GET request to httpbin.org/get and prints the data to stdout
func demoGet(done chan bool) {

	// Getting json data from the web
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// injecting some artificial delay
	n := rand.Intn(2000)
	time.Sleep(time.Duration(n) * time.Millisecond)

	// var data map[string]interface{}
	data := make(map[string]interface{})
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range data {
		fmt.Printf("%v: %v\n", k, v)
	}

	done <- true
}

func main() {

	// json formatted string
	zipCodesStr := `{"San Jose" : 84753,"Fairfax" : 73949,"Cole" : 49585}`

	// a decoder to read the data json formatted data
	decoder := json.NewDecoder(strings.NewReader(zipCodesStr))

	// memory to store the json data
	var zipCodes map[string]interface{}

	for {
		if err := decoder.Decode(&zipCodes); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Unexpected error in decoding json data")
		} else {
			for city, zip := range zipCodes {
				fmt.Printf("city: %v, zip:%v\n", city, zip)
			}
		}
	}

	done := make(chan bool)
	go demoGet(done)

	// Wait for the completion on the done channel
	// if the completion doesn't arrive in a second, timer will
	// elapse and program will exit. The random delay included
	// in demoGet controls whether the timer gets elapsed or
	// main gets the done notification before the timer expires.
	select {
	case <-time.After(time.Second):
		fmt.Println("Timer elapsed before completing the request")
		break
	case <-done:
		fmt.Println("Completed GET successfully")
		break
	}
}
