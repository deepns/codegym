package main

// importing single package
import (
	"errors"
	"fmt"
	"math"
)

// importing multiple packages

func main() {
	fmt.Println("Learning Basics")

	// using format verbs
	// %v - for the value
	// %T - to get the type
	fmt.Printf("pi=%v, type(pi)=%T\n", math.Pi, math.Pi)

	// creating a new error with errors package
	fmt.Println(errors.New("SampleError"))

	for i := 0.0; i < 3; i++ { // braces should end in the same line as the block
		fmt.Printf("i=%v, sqrt(i)=%v\n", i, math.Sqrt(i))
		fmt.Printf("i=%v, sqr(i)=%v\n", i, math.Pow(i, 2))
	}
}
