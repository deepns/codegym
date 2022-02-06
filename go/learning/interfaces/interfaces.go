package main

import (
	"fmt"
	"net"
)

// In languages like Java, interfaces are explicit contracts
// that lists set of methods to be implemented. Go interfaces
// are implicit.
//
// interface type defines a set of method signatures
// interface value can hold any value that implement those methods
// under the hood, interface ~= (value, type)
//
// A perfect starter example is the Stringer interface from "fmt" package
// Any type that has String() method implicitly implements the Stringer
// interface. fmt.Print.. internally calls the String() method on the
// interface value if the value is of type Stringer.
// See handleMethods of type "pp" in print.go
// https://github.com/golang/go/blob/master/src/fmt/print.go#L570
//
// // Stringer is implemented by any value that has a String method,
// // which defines the ``native'' format for that value.
// // The String method is used to print values passed as an operand
// // to any format that accepts a string or to an unformatted printer
// // such as Print.
// type Stringer interface {
// 	String() string
// }
//
// // GoStringer is implemented by any value that has a GoString method,
// // which defines the Go syntax for that value.
// // The GoString method is used to print values passed as an operand
// // to a %#v format.
// type GoStringer interface {
// 	GoString() string
// }
//
// In Go, it is idiomatic to name the interface as "method-er" that
// has a "method". For e.g. Stringer has String().	Reader() has Read()
// Formatter() has Format()
func Basics() {
	serverAddr := Addr{
		net.IPv4(127, 0, 0, 1),
		8080,
	}
	fmt.Printf("serverAddr: %v\n", serverAddr)
}

// Defining types that implements Stringer

type Addr struct {
	Ip   net.IP
	Port uint16
}

func (a Addr) String() string {
	return fmt.Sprintf("%v:%v", a.Ip, a.Port)
}
