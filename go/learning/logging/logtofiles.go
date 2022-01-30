package main

import (
	"fmt"
	"log"
	"os"
)

// LogToFiles shows an example of writing the logs to a file
func LogToFiles() {
	// The standard logs to stdout.
	// We can create custom loggers to write to in-memory buffers (e.g. bytes.Buffer)
	// or files, or any type that supports io.Writer interface.
	fmt.Println("====== Logging to files ======")

	// os.Create returns os.File which supports io.Writer
	logfile, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	fLogger := log.New(logfile, "[ F-LOG ] ", log.LstdFlags|log.Lmsgprefix)
	fLogger.Print("Created the log file")
}
