// Learning Slices
package main

import (
	"fmt"
	"sort"
)

// Basics of slices
//	defintion
//	append
// 	make vs new
//	iteration
//	sorting
//	splicing
func main() {
	fmt.Println("========== Learning Slices ==========")
	// Slices are declared similar to slice, except without
	// a size parameter. since no size is specified, there is
	// no memory allocated for this until the append time.
	// a slice doesn't store data by itself. it uses an array
	// underneath
	var aSliceInt []int

	fmt.Printf("aSliceInt == nil: %v\n", aSliceInt == nil)

	// expand slice with append
	// %#v shows the Go-syntax representation of the value
	aSliceInt = append(aSliceInt, 101)
	fmt.Printf("aSliceInt=%#v, aSliceInt=%T\n", aSliceInt, aSliceInt)

	// implicit declaration
	aSliceFloat := []float64{102.0, 101.1, 98.4, 103.4, 107.4, 100}
	fmt.Printf("aSliceFloat=%v, aSliceFloat=%T, len(aSliceFloat)=%v\n",
		aSliceFloat, aSliceFloat, len(aSliceFloat))

	// sort slice using built in sort functions
	sort.Float64s(aSliceFloat)
	fmt.Printf("After sorting, aSliceFloat=%v\n", aSliceFloat)

	// create a slice with make
	// make is used to allocate and initialize slice, map and channels.
	// slice take the length and initial capacity parameters
	aSliceUsingMake := make([]int, 0 /*length*/, 5 /*capacity*/)

	// both len() and cap() are built in methods
	fmt.Printf("aSliceUsingMake=%v, len(aSliceUsingMake)=%v, cap(aSliceUsingMake)=%v\n",
		aSliceUsingMake, len(aSliceUsingMake), cap(aSliceUsingMake))

	// note that the length of the slice was still 0, so we cannot
	// assign directly using the index. we have to use append() to
	// add to the slice
	// Add 5 elements to the slice
	for i := 0; i < 5; i++ {
		aSliceUsingMake = append(aSliceUsingMake, i)
	}

	// expand the slice after reaching the full capacity
	// capacity will be doubled.
	aSliceUsingMake = append(aSliceUsingMake, 6)

	// Now check the increase in capacity of the slice
	fmt.Printf("aSliceUsingMake=%v, len(aSliceUsingMake)=%v, cap(aSliceUsingMake)=%v\n",
		aSliceUsingMake, len(aSliceUsingMake), cap(aSliceUsingMake))

	// while make is specific to slice, map and channels,
	// new() is generic to all types. Note, it returns the
	// pointer to the allocated memory. have to see where
	// new is preferred and used commonly.
	aSliceUsingNew := new([]int)
	fmt.Printf("aSliceUsingNew=%T\n", aSliceUsingNew)

	*aSliceUsingNew = append(*aSliceUsingNew, 2048)
	fmt.Printf("aSliceUsingNew=%v, aSliceUsingNew=%T\n", aSliceUsingNew, aSliceUsingNew)

	// slices can be iterated using the range keyword
	// for slices, range returns index, value
	fmt.Println("Iterating a slice...")
	for i, v := range aSliceFloat {
		fmt.Printf("aSliceFloat[%v]=%v\n", i, v)
	}

	// I really the shorcuts in vscode to insert code patterns
	// for e.g slice have shortcuts for range, copy, last, append, sort and some more.
	sort.Slice(aSliceFloat, func(i, j int) bool {
		return aSliceFloat[i] > aSliceFloat[j]
	})
	fmt.Printf("After reverse sorting: aSliceFloat: %v\n", aSliceFloat)

	// reversing a slice
	for i, j := 0, len(aSliceFloat)-1; i < j; i, j = i+1, j-1 {
		aSliceFloat[i], aSliceFloat[j] = aSliceFloat[j], aSliceFloat[i]
	}

	// splicing a slice
	fmt.Printf("aSliceFloat[:2]: %v\n", aSliceFloat[:2])
	fmt.Printf("aSliceFloat[3:]: %v\n", aSliceFloat[3:])
	fmt.Printf("aSliceFloat[1:4]: %v\n", aSliceFloat[1:4])
	// negative indexes are not supported. Hail Python!
	// fmt.Printf("aSliceFloat[:-2]: %v\n", aSliceFloat[:-2])
	fmt.Printf("aSliceFloat[:]: %v\n", aSliceFloat[:])

}
