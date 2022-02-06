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
	"github.com/deepns/codegarage/go/learning/collections/structs"
)

func main() {
	arrays.Basics()
	maps.Basics()
	structs.Basics()
	structs.Methods()
	structs.Embed()
}
