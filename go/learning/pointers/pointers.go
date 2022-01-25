package main

/*
 * Trying out pointers in Go
 */

import (
	"fmt"
)

func main() {
	var anInt int
	// Declaring a pointer to an integer explicitly
	// it is good assign the pointer at declaration itself.
	// go static check cries foul otherwise
	var anIntPtr *int = &anInt
	*anIntPtr = 101

	fmt.Printf("anInt=%v, anIntPtr=%v, *anIntrPtr=%v\n", anInt, anIntPtr, *anIntPtr)

	// implicit pointer
	aFloat := 12.034
	aFloatPtr := &aFloat

	fmt.Printf("aFloat=%v, aFloatPtr=%v, *aFloatPtr=%v, type(aFloatPtr)=%T\n",
		aFloat, aFloatPtr, *aFloatPtr, aFloatPtr)

	// pointer to a slice (more on slices later)
	aSliceInts := []int{2, 4, 6, 8}
	aSliceIntsPtr := &aSliceInts

	// Note that %v for a pointer prints &<value pointed by the ptr>
	// below statement would print aSliceIntsPtr=&[2 4 6 8]
	// however for simple types like aFloatPtr, %v prints the address directly
	fmt.Printf("aSliceInts=%v, aSliceIntsPtr(%%v)=%v, aSliceIntsPtr(%%p)=%p, "+
		"*aSliceIntsPtr=%v, type(aSliceIntsPtr)=%T\n",
		aSliceInts, aSliceIntsPtr, aSliceIntsPtr, *aSliceIntsPtr, aSliceIntsPtr)

	// Just like C, variables passed to functions are copied if passed by value.
	// variables with high memory footprint are better passed as pointers
	AddToList(&aSliceInts, 10)
	fmt.Printf("After AddToList(&aSliceInts, 12), aSliceInts=%v\n", aSliceInts)

	// using the built in append
	aSliceInts = append(aSliceInts, 12)

	// Pointers become more interesting with structs
	// we can define methods that operate on struct pointers
	// with structs, go supports "." notation to access struct
	// fields with pointers as well.
}

func AddToList(list *[]int, anInt int) {
	*list = append(*list, anInt)
}
