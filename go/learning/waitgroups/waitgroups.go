// Exploring waitgroups from sync
// WaitGroup waits for a collection of go routines to finish.
// Callers calls Add() on the waitrgroup for each go routine
// which then calls Done() when the routine finishes.
package main

import (
	"fmt"
	"sync"
	"time"
)

func doJob(id int) {
	fmt.Println("Doing job", id)
	time.Sleep(2 * time.Second)
	fmt.Println("Done job", id)
}

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)

		go func() {
			// Done decrements the wg Counter by one
			defer wg.Done()
			doJob(i)
		}()
	}

	// Wait() blocks until wg counter is zero
	wg.Wait()

	// Go src seem to be using WaitGroup pretty commonly
	// check ServeCodec from net/rpc/server.go

}
