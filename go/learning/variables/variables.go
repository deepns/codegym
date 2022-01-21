package main

import "fmt"

// Names starting in uppercase are exported
var Exported = true

func main() {
	// declaring types explicitly
	// takes the default value for the type
	// default value for int is 0
	var aInt int

	// declaring and assigning the value
	var aIntAssigned int = 31

	// strings
	var aName string = "Foo"

	// implicit typing
	var earthIsFlat = false

	// implicit typing with assignment
	rate := 3.5

	// multiple variables can be declared at the same time
	tic, tac, toe := true, false, true

	const aConstInt int = 101
	const aConstString string = "FOO"
	const aConstBool = false
	const pi = 3.14

	// changing constant will be caught in compilation itself
	// aConstInt = 24

	fmt.Printf("pi=%v, type(pi)=%T\n", rate, rate)

	// unused variables are not allowed in Go.
	fmt.Printf("aInt=%v, type(aInt)=%T\n", aInt, aInt)
	fmt.Printf("aIntAssigned=%v, type(aIntAssigned)=%T\n", aIntAssigned, aIntAssigned)
	fmt.Printf("aName=%v, type(aName)=%T\n", aName, aName)
	fmt.Printf("earthIsFlat=%v, type(earthIsFlat)=%T\n", earthIsFlat, earthIsFlat)
	fmt.Printf("tic=%v, tac=%v, toe=%v\n", tic, tac, toe)

	fmt.Printf("aConstInt=%v, type(aConstInt)=%T\n", aConstInt, aConstInt)
	fmt.Printf("aConstString=%v, type(aConstString)=%T\n", aConstString, aConstString)
	fmt.Printf("aConstBool=%v, type(aConstBool)=%T\n", aConstBool, aConstBool)
	fmt.Printf("pi=%v, type(pi)=%T\n", pi, pi)
}
