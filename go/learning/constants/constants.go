package main

import (
	"fmt"
	"math"
)

func main() {
	// constants can be of strings, character boolean or numeric values
	// constants can be typed or untyped

	// string constants
	const proto = "https://"
	const homepage = proto + "www.google.com"

	// there is no char type in go.
	// rune is used as the type for chars
	const aCharConstant = 'A'

	// a float constant
	const pi = 3.14

	// const can appear wherever var can.
	// const can be grouped too.
	const (
		MONDAY  = 1
		TUESDAY = MONDAY + 1 // const expressions can included arithmetic too
	)

	// a boolean constant
	const isEnabled = true

	fmt.Printf("homepage: %v, type:%[1]T\n", homepage)
	fmt.Printf("aCharConstant: %v, type:%[1]T\n", aCharConstant)
	fmt.Printf("pi: %v, type:%[1]T\n", pi)
	fmt.Printf("TUESDAY: %v, type:%[1]T\n", TUESDAY)
	fmt.Printf("isEnabled: %v, type:%[1]T\n", isEnabled)

	const X, Y = -4, -8.0

	// with untyped constants, types can be given based
	// on the context. for e.g. math.Abs() takes a float64,
	// so X an untyped int can be used as float64 there.
	fmt.Printf("math.Abs(X): %v\n", math.Abs(X))
	fmt.Printf("math.Abs(Y): %v\n", math.Abs(Y))

	// a typed constant
	const Z int = 12

	// Since Z is typed, it needs to be converted to float64
	// before using it in a context where float64 is needed.
	fmt.Printf("math.Abs(float64(Z)): %v\n", math.Abs(float64(Z)))
}
