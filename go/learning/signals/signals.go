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

	// register multiple signals
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2)

	// ignore selective signals
	signal.Ignore(syscall.SIGABRT)

	// wait until signal is received
	fmt.Println("Waiting for signal. Press Ctrl-C")
	sig := <-sigs
	fmt.Println("Received", sig, "signal")

	done := make(chan bool)
	go func() {
		fmt.Println("Waiting for signal in a go routine")
		sig := <-sigs
		fmt.Println("Received", sig, "signal")
		done <- true
	}()

	<-done

	// Not many examples of signal notification in the standard library
	// some examples in other repos
	// https://github.com/kubernetes/kubernetes/blob/eb43b41cfd59adfb0ee88e34f5967a8cf6ed0c9b/staging/src/k8s.io/kubectl/pkg/cmd/profiling.go#L72
	// from kubectl - https://github.com/kubernetes/kubectl/blob/652881798563c00c1895ded6ced819030bfaa4d7/pkg/util/interrupt/interrupt.go#L90
}
