package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// DoSomething waits for a random duration between 0 and 10 seconds, and then
// prints "Done!" to the console. If the context passed to the function is cancelled
// before the sleep duration has elapsed, the function will print "Timed out!" to
// the console instead.
//
// The function uses a timer to wait for the sleep duration, and a context to
// handle cancellations. If the timer fires before the context is cancelled, the
// function will print "Done!" and return. If the context is cancelled before the
// timer fires, the function will print "Timed out!" and return.
//
// Note that the function uses a random sleep duration to simulate a long-running
// operation that cannot be cancelled directly.
func DoSomething(ctx context.Context) {
	rand.Seed(int64(time.Now().UnixNano()))
	sleepTime := time.Second * time.Duration(rand.Intn(10))
	log.Printf("Waiting for %v", sleepTime)

	// Why NewTimer() instead of using time.After()?
	// With time.After(), underlying timer is not garbage collected until the
	// timer fires. In this case, if the parent context sends a cancellation signal
	// before the timer fires, the allocated timer will be leaked.
	// Not a big deal for this example since this is short lived.
	// Good to be aware of such gaps.
	t := time.NewTimer(sleepTime)
	defer t.Stop()

	select {
	case <-t.C:
		fmt.Println("Done!")
	case <-ctx.Done():
		fmt.Println("Timed out!")
	}
}

func main() {
	// WithTimeout takes a context and returns derived context (with a
	// timer to be fired upon timeout) and a cancelFunc. The parent context
	// here is provided by context.Background() which returns a non-nil, empty context.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	// cancel the timer and release resources if DoSomething completes before timeout elapses
	defer cancel()

	// If DoSomething takes longer than 5 seconds, a cancellation signal will be sent
	DoSomething(ctx)

	// Some rules of thumb to follow when using contexts
	// - incoming requests to a server should create a context, outgoing calls to a server
	//	 should accept a context
	// - chain of function calls between them must propagate the context, optionally
	//   deriving new context from the parent
	// - context should be passed explicitly, typically as the first argument to a function.
	//	 Do not store the context inside a struct type
	// - use context values only for request scoped data that transits processes and API

	// Example use case
	// a web server that handles user requests. Some of these requests may take a long time
	// to process, such as requests that involve complex database queries or remote API calls.
	// To handle these long-running requests, we might use a separate Goroutine to perform the work,
	// so that the main Goroutine can continue to handle other requests.

	// However, we also want to be able to cancel long-running requests if the user cancels the request
	// or the connection is closed. To do this, we could use the context package to create a context
	// for each request, and pass that context to the Goroutine handling the request. If the request
	// is cancelled, the context's Done() channel will be closed, and the Goroutine can exit gracefully.

	// Some examples from the standard library using contexts
	// - https://pkg.go.dev/net/http#NewRequestWithContext
	// - https://pkg.go.dev/database/sql#DB.QueryContext

	// Outside the standard library, gRPC also uses context heavily
}
