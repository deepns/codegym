// Learning errors in Go
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	// it is idiomatic in Go to communicate errors via explicit
	// return value. When returning multiple values, errors are
	// the last return value by convention.
	// for e.g. os.Open() returns a FILE pointer and error
	// if file can be opened successfully, err will be nil.
	// opening an non existing file here.
	_, err := os.Open("a-missing-file")
	if err != nil {
		// opening a missing file generates fs.PathError
		// not using log.Fatal() since I want the execution to
		// continue.
		fmt.Printf("err: %v, type:%T\n", err, err)
	}

	// errors are one of the simplest packages in Go.
	// types can implement the built-in interface error
	// to create custom error types.
	// Create new errors with errors.New, which creates an error
	// of type errors.errorString.

	var port int
	_, err = fmt.Scanf("%d", &port)
	if err != nil {
		log.Fatal("Failed to scan port number")
	}

	err = checkPort(port)
	fmt.Printf("err: %v, type(err): %[1]v\n", err)
}

func checkPort(port int) error {
	if port == 0 {
		return errors.New("port 0 is reserved")
	}
	return nil
}
