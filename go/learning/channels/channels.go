// Exploring channels
//
// channels - built in go types to do buffered and unbuffered bidrectional communication
// like slice and maps, channels are typed, and created with make
// send to channel: "chan <- val", blocks until is ready to take in new data
// recv from channel: "val <-chan", blocks until chan has some new data to read
// channels support range iteration too.
//	for m := range aChannel {
//		// loops until the channel is closed
//	}
package main

import (
	"fmt"
	"time"
)

func main() {
	nGoRoutines := 10

	// an unbuffered channel
	done := make(chan int)

	for i := 0; i < nGoRoutines; i++ {
		go func(id int, done chan int) {
			time.Sleep(1 * time.Second)
			done <- id
		}(i, done)
	}

	for i := 0; i < nGoRoutines; i++ {
		id := <-done
		fmt.Printf("go routine %d is done\n", id)
	}

	fmt.Println("Making an unbuffered channel")

	// messages is a buffered channel, takes up to 3 messages in the buffer
	messages := make(chan string, 3)

	// like slices and maps, channels also support length and capacity
	fmt.Printf("len(messages): %v\n", len(messages))
	fmt.Printf("cap(messages): %v\n", cap(messages))

	messages <- "get"
	messages <- "set"
	messages <- "go"

	fmt.Printf("After sending three messages: len(messages): %v\n", len(messages))

	// launch a go routine to read from the channel
	go func() {
		// sleeping some time intentionally so caller would wait until messages
		// channel can take in new data
		time.Sleep(5 * time.Second)
		for m := range messages {
			fmt.Printf("m: %v\n", m)
		}
	}()

	// Sending any more would block
	messages <- "I know you will block"
	messages <- "This won't block"

	// closing the channel
	// channels don't require to be closed usually. If the receiver is iterating
	// on the channel with a for .. range loop, then closing the channel helps to
	// indicate to the receiver to stop iterating
	close(messages)

	// check whether a channel is closed or not
	_, ok := <-messages
	if !ok {
		fmt.Printf("%v is closed\n", messages)
	}
}
