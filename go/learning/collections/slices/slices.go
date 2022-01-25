package slices

import (
	"fmt"
	"sort"
)

func Learn() {
	fmt.Println("========== Learning Slices ==========")
	// Slices are declared similar to slice, except without
	// a size parameter. since no size is specified, there is
	// no memory allocated for this until the append time.
	var aSliceInt []int

	// expand slice with append
	// %#v shows the Go-syntax representation of the value
	aSliceInt = append(aSliceInt, 101)
	fmt.Printf("aSliceInt=%#v, aSliceInt=%T\n", aSliceInt, aSliceInt)

	// implicit declaration
	aSliceFloat := []float64{102.0, 101.1, 98.4}
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
}
