// Exploring the init function
// init() is a special function, much like main
// init() is called after packages are imported and initialized, and variables
// in the local package are initialized.
// imported packages can have their own init functions.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var inited bool

// A very simple use of init.
// setting up the prog state at the start.
// init() makes much more sense when used in a package that is not main
func init() {
	fmt.Printf("inited: %v\n", inited)
	rand.Seed(time.Now().UnixNano())
	inited = true
}

// init() functions can be defined multiple times
// they are called in the order in which they are encountered
// by the compiler. But what is the point of having multiple
// init functions? If a package has multiple files, and each file
// has their own init functions, how does the ordering work?
func init() {
	fmt.Println("Second init")
}

func main() {
	fmt.Printf("inited: %v\n", inited)
	fmt.Printf("rand.Intn(1000): %v\n", rand.Intn(1000))

	// Some examples of init() in go src
	// https://github.com/golang/go/blob/master/src/crypto/md5/md5.go
	// uses the init() for registration

	// likewise, https://github.com/golang/go/blob/master/src/image/png/reader.go
	// also uses init() for registration

	// here is another one.
	// https://github.com/golang/go/blob/master/src/crypto/rand/rand_unix.go

	// when packages are imported for their side effect, they are imported
	// with a blank identifier
	// see comments https://github.com/golang/go/blob/master/src/image/image.go
}
