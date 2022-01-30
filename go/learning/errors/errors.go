// Learning errors in Go
package main

import (
	"errors"
	"fmt"
	"net/http"
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

	// errors can be printed by calling Error() on the error type
	// or by directing printing them with fmt (fmt.Print internally
	// calls Error())
	if err != nil {
		// opening a missing file generates fs.PathError
		// not using log.Fatal() since I want the execution to
		// continue.
		fmt.Printf("err: %v, type:%T\n", err, err)

		// Get error details with Error
		fmt.Printf("err.Error(): %v\n", err.Error())
	}

	// errors are one of the simplest packages in Go.
	// types can implement the built-in interface error
	// to create custom error types.
	// New errors can be created with
	//	fmt.Errorf
	//	errors.New()
	// These errors of type errors.errorString.

	var someNumber int
	fmt.Print("Enter a number between 0-10: ")
	fmt.Scanf("%d", &someNumber)

	err = nil
	if someNumber < 0 {
		err = errors.New("number less than 0")
	} else if someNumber > 10 {
		err = fmt.Errorf("number %v greater than 10", someNumber)
	}
	fmt.Printf("err: %v, type(err): %[1]T\n", err)

	// Custom errors can be created by adding Error() to custom types
	// InvalidPortError is one such example.
	port := -1
	if port < 0 || port > 65535 {
		err = InvalidPortError(port)
		fmt.Printf("err: %v, type(err): %[1]T\n", err)
	}

	someInvalidAddress := "http://osdogsigsgosigs.dxgs"
	_, err = http.Get(someInvalidAddress)
	fmt.Printf("err: %v, type(err): %[1]T\n", err)

	// errors can be extended with additional information too.
	// here is an example from DNSError
	// DNSError represents a DNS lookup error.
	// type DNSError struct {
	// 	Err         string // description of the error
	// 	Name        string // name looked for
	// 	Server      string // server used
	// 	IsTimeout   bool   // if true, timed out; not all timeouts set this
	// 	IsTemporary bool   // if true, error is temporary; not all errors set this
	// 	IsNotFound  bool   // if true, host could not be found
	// }
}

// InvalidPortError is used when the port number is beyond the range 0-65535
type InvalidPortError int

func (p InvalidPortError) Error() string {
	return fmt.Sprintf("invalid port number: %v", int(p))
}
