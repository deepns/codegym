package functions

import "fmt"

// func syntax
// func funcName(args) return {...}

// FuncWithNoArgNoReturn takes no arg, returns nothing
func FuncWithNoArgNoReturn() {
	fmt.Println("FuncWithNoArgNoReturn")
}

// Add returns the sum of two numbers
// function with args and a return type
func Add(a int, b int) int {
	return a + b
}

// Sub returns the sum of two numbers
// function with args and a return type
func Sub(a int, b int) int {
	return a - b
}

// GetFunc returns a function for the given operator
func GetFunc(op string) func(int, int) int {
	if op == "Sub" {
		return Sub
	} else {
		return Add
	}
}

// Double is an example for FuncWithPointerArgs
// doubles given the number in place
func Double(ptrA *int) {
	*ptrA = *ptrA * 2
}
