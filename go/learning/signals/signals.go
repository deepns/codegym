// Exploring signal handling
// signals handled using os and os/signal packages
// https://pkg.go.dev/os#Signal defines the interface for Signal
// only kill and interrupt signal values are guaranteed to be present in
// the os package.
// https://pkg.go.dev/syscall#Signal implements the os.Signal interface

// signal registration and catching is handled using a buffered channel
// register for signal notification through signal.Notify
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// package signal does not block sending to the channel that is
	// registered to receive the signal. so the channel must be buffered.
	// the buffer size is application dependent.
	sigs := make(chan os.Signal, 1)

	// register to receive SIGINT
	// can also use os.Interrupt in place of syscall.SIGINT. They are equal.
	signal.Notify(sigs, syscall.SIGINT)

	// wait until signal is received
	fmt.Println("Waiting for signal. Press Ctrl-C")
	sig := <-sigs
	fmt.Println("Received", sig, "signal")

	// TODO
	// - [ ] handle signal in a separate go routine
	// - [ ] ignore signals
	// - [ ] handle additional signals (see syscall package)
	// - [ ] find examples from package
}
