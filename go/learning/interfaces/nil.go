package main

import "fmt"

// Exploring nil interfaces

type Tracer interface {
	trace(msg string)
}

type Packet struct {
	Type uint16
	Data []byte
}

func (p *Packet) trace(msg string) {
	if p == nil {
		fmt.Println("receiver is nil")
		return
	}
	fmt.Println("packet trace:", msg)
}

// NilInterface explores the usage of nil interfaces, calling
// interface methods with nil receivers
func NilInterface() {
	var t Tracer
	// %T verb usually shows the type of the value
	// However, that doesn't work with interfaces
	// There is no inhererent type or value for an interface
	// until the interface holds a value underneath.
	// The below statement would print
	// t: <nil>, type(t):<nil>
	// calling methods on a nil interface is a runtime error
	fmt.Printf("t: %v, type(t):%T\n", t, t)

	var p *Packet
	fmt.Printf("p: %v, type(p):%T\n", p, p)
	fmt.Printf("(p == nil):%v\n", (p == nil))

	// Given that p is nil, calling a method on a nil (null ptr)
	// would usually trigger NullPointer error on most languages.
	// Go handles it differently. A function receiver can be nil.
	// It is up to the implementation to handle it.
	p.trace("nil ptr")

	// This allows interface to call methods even when the
	// underlying value of an interface is nil
	t = p

	// Now that t's underlying value is p, %T on the interface
	// would print the type of the underlying value
	// This would print t: <nil>, type(t):*main.Packet
	fmt.Printf("t: %v, type(t):%T\n", t, t)
	t.trace("from interface")
}
