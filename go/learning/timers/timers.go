// Exploring timers from time package
// timer represents single event in the future
// when a timer expires, current time is sent on the timer's channel
//
// Note that a timer must be create with NewTimer or AfterFunc.
// A timer has a channel that is exposed and a runtimeTimer that
// is private to the Timer. Comments from the Timer struct.
//
// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C,
// unless the Timer was created by AfterFunc.
// A Timer must be created with NewTimer or AfterFunc.
// type Timer struct {
// 	C <-chan Time
// 	r runtimeTimer
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	// create a timer to be fired after 5 seconds
	t1 := time.NewTimer(5 * time.Second)

	// blocks until t1.C has something to read.
	curTime := <-t1.C

	fmt.Println("Timer expired at", curTime.UTC())

	// Resetting the timer
	t1.Reset(5 * time.Second)

	// Stop timer after some time.
	// AfterFunc fires off another timer which executes the given func
	// when the timer expires.
	time.AfterFunc(2*time.Second, func() {
		stop := t1.Stop()
		if stop {
			fmt.Println("Timer stopped after 2 seconds")
		} else {
			// Drain the channel if timer can't be stopped
			<-t1.C
		}
	})

	// giving some time for the timer to stop
	time.Sleep(3 * time.Second)

	// Here is an example from net/http/client.go that makes use of timer
	// https://github.com/golang/go/blob/master/src/net/http/client.go#L394
	// timer := time.NewTimer(time.Until(deadline))
	// var timedOut atomicBool

	// go func() {
	// 	select {
	// 	case <-initialReqCancel:
	// 		doCancel()
	// 		timer.Stop()
	// 	case <-timer.C:
	// 		timedOut.setTrue()
	// 		doCancel()
	// 	case <-stopTimerCh:
	// 		timer.Stop()
	// 	}
	// }()
}
