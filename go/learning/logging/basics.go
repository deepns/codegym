package main

import (
	"fmt"
	"log"
	"os"
)

// Basics shows some examples of getting started with log package
// Use log.Print to send log messages to standard out
// Change log flags
// Add custom prefix
func Basics() {

	fmt.Println("====== Some Basics ======")

	// log has a predefined logger that writes to stdout.
	// They are accessed through Print/Fatal/Panic functions

	// log.Print/Println/Printf prints
	// <date> <time> message
	log.Println("starting main()")

	// log constants
	// they control what fields to print along with given message
	// date and time flags are set by default.
	// NOTE: If the message doesn't include a newline at the end,
	// log.Print()/Printf() automatically adds a newline when printing
	// the logss
	log.Printf("Default flags: %#x", log.Flags())

	// lets add the filename and linenumber and also print the time in UTC format
	// NOTE: The order in which these info shows in the log is not the
	// same in which they are defined. It is determined by the logger.
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
	log.Printf("New Flags: %#x", log.Flags())

	log.SetPrefix("[ INFO ] ")
	log.SetFlags(log.Flags() | log.Lmsgprefix)
	log.Println("Added the prefix")

	// Default logger writes to stderr
	fmt.Println("Default writer:", log.Default().Writer(), "os.Stderr:", os.Stderr)

	// log.Fatal, log.Fatalln, log.FatalF is equivalent to calling
	// the corresponding Print functions and os.Exit(1)

	// log.Panic, log.Panicln, log.PanicF is equivalent to calling
	// the corresponding Print functions and panic()

	// There is no predefined LOG LEVELS e.g. ERROR/WARN/INFO/DEBUG/FATAL
	// I guess that is left to the application to manage their custom
	// loggers.
}
