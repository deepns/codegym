package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// There is no pre-defined log levels in Go. It is up to
// the application to define the log levels and create
// custom loggers for each of the log level. Here is an
// example.
// If we define this as a separate package (for e.g. logs)
// other modules can then import and use the defined loggers.
// e.g. logs.Info.Print(...)
// logs.Error.Fatal(...)

// io.Discard is equivalent to /dev/null. All writes to
// that writer succeeds.

var (
	// Error is used to log fatal errors
	Error = log.New(io.Discard, "[ ERROR ] ", log.LstdFlags)

	// Info is used to log notable events
	Info = log.New(io.Discard, "[ INFO ] ", log.LstdFlags)

	// Debug is used to log information useful for debugging
	Debug = log.New(io.Discard, "[ DEBUG ] ", log.LstdFlags)
)

// CustomLogging shows how we can define custom loggers to handle
// different logging levels based on the application needs.
func CustomLogging() {
	fmt.Println("====== Custom logging with different levels ======")
	// Lets send the error logs to stderr, info logs to
	// stdout. Not caring about debug logs, so leave it io.Discard
	Error.SetOutput(os.Stderr)
	Info.SetOutput(os.Stdout)

	Error.Print("This will be sent to stderr")
	Info.Print("This will be sent to stdout")
	Debug.Print("This will be discarded")
}
