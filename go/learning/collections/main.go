package main

/*
 * Learning collection types in Go.
 * 	- Arrays
 *  - Slices
 *  - Maps
 *  - Structs
 * Going with a slightly different approach this time
 * Instead of having main() for example, I have defined
 * each in their own package
 */

import (
	"github.com/deepns/codegarage/go/learning/collections/arrays"
	"github.com/deepns/codegarage/go/learning/collections/maps"
)

func main() {
	arrays.Basics()
	maps.Basics()
}
