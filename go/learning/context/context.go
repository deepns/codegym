package main

import (
	"context"
	"fmt"
	"time"
)

func doSomeWork(ctx context.Context) {
	time.Sleep(time.Second * 2)
	select {
	case <-ctx.Done():
		fmt.Println("Work canceled!")
	default:
		fmt.Println("Work done!")
	}
	// TODO
	// [ ] Send completion on the context?
	// main() still waits even after doSomeWork completes. Why?
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()

	go doSomeWork(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Work timed out")
	}
}
