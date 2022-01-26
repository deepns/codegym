package structs

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

func LearnStructMethods() {
	localServer := Connection{
		"localhost",
		net.IPv4(127, 0, 0, 1),
		8080,
	}

	// Since Connection has String(), that will be used instead of
	// standard print which will print all fields of the struct
	fmt.Println(localServer)
	fmt.Printf("localServer: %v\n", localServer)

	// print the go representation of the type
	fmt.Printf("localServer: %#v\n", localServer)
}
