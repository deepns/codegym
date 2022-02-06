package main

import (
	"fmt"
	"net"
)

// Learning struct methods

type Connection struct {
	Name   string
	IPAddr net.IP
	Port   int
}

// To add methods to struct, add the struct var and type
// between func...functionName()
// String() is the equivalent of __str__ in Python
// String() returns a string representation of Connection
func (c Connection) String() string {
	return fmt.Sprintf("%v:%v", c.IPAddr, c.Port)
}

// IsPrivate returns true if the connection belongs to a private IP
func (c Connection) IsPrivate() bool {
	return c.IPAddr.IsPrivate()
}

// Methods create a struct type, and calls the methods
// of that struct
func Methods() {
	serverConn := Connection{
		"localhost",
		net.IPv4(127, 0, 0, 1),
		8080,
	}

	// Since Connection has String(), that will be used instead of
	// standard print which will print all fields of the struct
	fmt.Println(serverConn)
	fmt.Printf("serverConn: %v, isPrivate:%v\n", serverConn, serverConn.IsPrivate())

	// print the go representation of the type
	fmt.Printf("serverConn: %#v\n", serverConn)
}
