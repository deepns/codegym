// Exploring select
// select {..} is similar to switch {...} where all the cases refer to communication operations
// select blocks until one of the communication cases is ready
// if a default: case is specified, then select runs the default case without blocking when none
// of the communication caes is ready
// the communication can either be send or receive
// works very similar to select in C
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// randOdd generates a random odd number on the channel until it is
// asked to quit in the quit channel
func RandOdd(c chan int, quit chan bool) {
	for {
		var n int
		for {
			n = rand.Intn(1000)
			if n%2 == 1 {
				break
			}
		}

		select {
		case c <- n:
			// blocks until n is sent to c
		case <-quit:
			// blocks until quit has something to read
			return
		}
	}

}

func main() {
	randChan := make(chan int)
	quit := make(chan bool)

	go func() {
		ticker := time.Tick(time.Second)
		ticks := 0

		for i := range ticker {
			x := <-randChan
			fmt.Printf("%v: x: %v\n", i, x)
			ticks += 1
			if ticks >= 10 {
				break
			}
		}
		
	}()

	RandOdd(randChan, quit)
}
