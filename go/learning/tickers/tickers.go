// Exploring tickers from time package
// Ticker fires off events in the future at defined intervals
package main

import (
	"fmt"
	"time"
)

func main() {

	// Tick every 500 ms.
	tick := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			// run until either the ticker channel is open or
			// done channel receives some data
			select {
			case <-tick.C:
				fmt.Println("Tick.")
			case <-done:
				fmt.Println("STOPPED!")
				return
			}
		}
	}()

	// Stop the ticker after some time
	time.AfterFunc(2 * time.Second, func() {
		tick.Stop()
		done <- true
	})

	// Wait for the ticker to stop.
	time.Sleep(3 * time.Second)
}