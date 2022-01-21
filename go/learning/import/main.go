package main

// importing another package from this module
// added a package named "functions" at the path
// go/learning/functions/. Combine the module name
// (defined in go.mod file at the top of the workspace)

import (
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
