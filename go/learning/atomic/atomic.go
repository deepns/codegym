// Exploring the atomic package
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// track the number of hits
	var nHits uint64

	// creating few thousand worker routines to update the counter in parallel
	// main needs to wait for the go routines to complete.
	// can use a channel where each go routine post its completion
	// or can use a wait group
	var wg sync.WaitGroup

	nWorkers := 10000
	fmt.Printf("Launching %v workers\n", nWorkers)
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&nHits, 1)
			wg.Done()
		}()
	}

	// wait until waitgroup is empty
	wg.Wait()
	fmt.Printf("nHits: %v\n", nHits)

	// first go routine to run will reset the hit counter
	// using compareandswap
	// probably not a good example of the CompareAndSwap usage. :(
	resetHitCounterWorker := func(id int) {
		swapped := atomic.CompareAndSwapUint64(&nHits, uint64(nWorkers), 0)
		fmt.Printf("worker: %v, swapped: %v\n", id, swapped)
		wg.Done()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go resetHitCounterWorker(i)
	}
	wg.Wait()

	// The example https://pkg.go.dev/sync/atomic#example-Value-ReadMostly in the docs
	// is pretty neat. Implementation of copy-on-write idiom for frequently read but infrequently
	// updated data. Data read without synchronization, and update using atomic Load (old value),
	// create new value, Store(new value) into atomic.Value

	// transport of http client uses atomic.Value to update the protocol defaults
	// func (t *Transport) onceSetNextProtoDefaults() {
	// t.tlsNextProtoWasNil = (t.TLSNextProto == nil)
	// if strings.Contains(os.Getenv("GODEBUG"), "http2client=0") {
	// 	return
	// }

	// If they've already configured http2 with
	// golang.org/x/net/http2 instead of the bundled copy, try to
	// get at its http2.Transport value (via the "https"
	// altproto map) so we can call CloseIdleConnections on it if
	// requested. (Issue 22891)
	// altProto, _ := t.altProto.Load().(map[string]RoundTripper)
	// if rv := reflect.ValueOf(altProto["https"]); rv.IsValid() && rv.Type().Kind() == reflect.Struct && rv.Type().NumField() == 1 {
	// 	if v := rv.Field(0); v.CanInterface() {
	// 		if h2i, ok := v.Interface().(h2Transport); ok {
	// 			t.h2transport = h2i
	// 			return
	// 		}
	// 	}
	// }
}
