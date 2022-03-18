// Exploring channels with directions
// Channels are bidirectional by default. They can be specified as send-only
// or recv only when used as function arguments or return values.
// Sending to a recv only channel or reading from a send-only channel
// is a compilation error.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Channels used to get notifications for a signal is
	// passed as send-only channel. See os/signal.Notify
	sigs := make(chan os.Signal, 1)
	// Registering to receive SIGINT on channel sigs
	// signal.Notify takes the channel as send only, so it
	// can't make any reads on that.
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		s := <-sigs
		fmt.Println("Received", s)
	}()

	// Channel returned by time.After is a recv only channel.
	t := time.After(5 * time.Second)
	// sending to a recv only channel is a compilation error.
	// t <- time.Now()
	fmt.Println("Waiting for the timer to expire")
	<-t
	fmt.Println("Timer expired.")
}
