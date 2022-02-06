package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Zip    int
	price  int
}

func (a Address) String() string {
	return fmt.Sprintf("%v %v - %v %v", a.Street, a.City, a.State, a.Zip)
}

// Another struct can be embedded by simply placing the struct
// name. Fields of the embedded struct can be accessed directly
// or through the embedded type. e.g. Property.Street and
// Property.Addr.Street are both valid.

type Property struct {
	Address
	IsResidential bool
}

// When embedding structs, both the fields and methods come along
// In this example, if Property() doesn't have its Stringer defined,
// Address's String() would be called. Some behavior to watch out
// for when using embedded structs
func (p Property) String() string {
	return fmt.Sprintf("%v, isResidential:%v", p.Address, p.IsResidential)
}

// Embed explores struct embedding
//	embedding works somewhat like inheritance
//	embed one struct in another
//	access exposed and private fields and methods of the
//	embedded struct
func Embed() {
	p1 := Property{
		Address{
			"111 Pen Ave",
			"Was",
			"DC",
			23233,
			250000,
		},
		false,
	}

	fmt.Printf("p1: %+v\n", p1)
	fmt.Printf("p1.Addr.City: %v\n", p1.Address.City)

	// Both exported and unexported fields of the embedded
	// type can be accessed
	fmt.Printf("p1.City: %v\n", p1.City)
	fmt.Printf("p1.price: %v\n", p1.price)

	fmt.Printf("p1.Address: %v\n", p1.Address)

	// some practical examples of embedding
	// ReadWriter in bufio -> https://github.com/golang/go/blob/2ebe77a2fda1ee9ff6fd9a3e08933ad1ebaea039/src/bufio/bufio.go#L775
	// lruSessionCache in crypto -> https://github.com/golang/go/blob/master/src/crypto/tls/common.go#L1388
	// type lruSessionCache struct {
	// 	sync.Mutex

	// 	m        map[string]*list.Element
	// 	q        *list.List
	// 	capacity int
	// }

	// Put/Get in lruSessionCache uses the mutex
	// calls Lock() / Unlock()

	// for e.g
	// makes the code more readable

	// func (c *lruSessionCache) Get(sessionKey string) (*ClientSessionState, bool) {
	// 	c.Lock()
	// 	defer c.Unlock()

	// 	if elem, ok := c.m[sessionKey]; ok {
	// 		c.q.MoveToFront(elem)
	// 		return elem.Value.(*lruSessionCacheEntry).state, true
	// 	}
	// 	return nil, false
	// }
}
