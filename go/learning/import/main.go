package main

// importing another package from this module
// added a package named "functions" at the path
// go/learning/functions/. Combine the module name
// (defined in go.mod file at the top of the workspace)

import (
	// go doesn't like packages imported but not used. There are some exceptions
	// where a package is imported (for init as a side effect) but not used explicitly.
	// in such case, we can suppress the error with the help of a blank identifier
	// https://go.dev/doc/effective_go#blank_import
	_ "flag"
	"fmt"
	"math/rand"

	"github.com/deepns/codegarage/go/learning/functions"
)

func main() {
	functions.FuncWithNoArgNoReturn()
	x, y := rand.Intn(94943), rand.Intn(1024)
	fmt.Printf("functions.Add(%v, %v)=%v\n", x, y, functions.Add(x, y))
	fmt.Printf("functions.Mod(%v, %v)=%v\n", x, y, functions.Mod(x, y))
	fmt.Printf("functions.GetFunc(\"Add\")(%v, %v)=%v\n", x, y, functions.GetFunc("Add")(x, y))
	fmt.Printf("y=%v,", y)
	// passing by pointer
	functions.Double(&y)
	fmt.Printf("After Double(&y)=%v\n", y)
}
