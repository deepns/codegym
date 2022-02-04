// Learning about variadic functions
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// variadic functions are similar to arbitrary args in python
// think of *args, **kwargs
// Does kwargs exist in go? Need to find out.
// func f(arg ...type) {} // to define a function with variadic input
// ...val // when passing a slice value to a variadic function

// sumInts returns sum of arbitrary number of integers
func sumInts(nums ...int) int {
	sum := 0

	// type of the arg with variadic inputs is a slice
	// of that type
	fmt.Printf("nums: %v, %[1]T\n", nums)
	for _, n := range nums {
		sum += n

	}
	return sum
}

// Another example for variadic functions
// Join concatenates the given strings with the rune
func Join(r rune, words ...string) string {
	sentence := ""
	for i, word := range words {
		sentence += word
		if i < len(words)-1 {
			sentence += string(r)
		}
	}
	return sentence
}

// DoubleUp doubles up all the given values
func DoubleUp(nums ...int) {
	// if a slice is unpacked at the caller, it is passed
	// by reference. so address of nums and the slice at
	// the caller will be the same.
	fmt.Printf("&nums=%p\n", nums)
	for i := range nums {
		nums[i] *= 2
	}
}

func main() {

	// making simple calls with variable number of parameters
	fmt.Printf("sumInts(10, -20): %v\n", sumInts(10, -20))
	fmt.Printf("sumInts(10, 20, 30): %v\n", sumInts(10, 20, 30))

	// Passing a slice to a variadic function

	nums := make([]int, 10)

	// Generating some random numbers
	// Not necessary for this example. but just having fun
	// Go's math/rand need to be seeded, otherwise rand results
	// are deterministic.
	rand.Seed(time.Now().Unix())
	for i := range nums {
		nums[i] = rand.Intn(100)
	}

	fmt.Printf("nums: %v\n", nums)
	// unpack a slice to a variadic function with ... notation
	fmt.Printf("sumInts(nums): %v\n", sumInts(nums...))

	// Using a variadic function that takes strings
	fmt.Printf("Join(',', \"hello\", \"world\"): %v\n", Join(',', "hello", "world"))
	names := []string{"alice", "bob", "mike"}
	fmt.Printf("Join(\"|\", names...): %v\n", Join('|', names...))

	// Want to see how a slice is unpacked when passed to variadic function
	// Does it make a copy of the elements or use the underlying array?
	somePrimes := []int{1, 2, 3, 5, 7}
	fmt.Printf("&somePrimes: %p\n", somePrimes)
	DoubleUp(somePrimes...)

	// Ok. the slice is indeed passed by reference in this case.
	fmt.Println("After doubling:")
	fmt.Printf("somePrimes: %v\n", somePrimes)

	// how about these ones? these are not accepted syntax
	// DoubleUp(somePrimes..., 9)
	// DoubleUp(11, somePrimes...)
}
