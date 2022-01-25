package arrays

import (
	"fmt"
)

func Learn() {
	fmt.Println("========== Learning Arrays ==========")

	// arrays are fixed size homogenous containers
	// once allocated, arrays cannot be shrunk and expanded

	var anIntArray [3]int
	anIntArray[0] = 101
	anIntArray[1] = 222
	anIntArray[2] = 103
	// simply overwriting
	anIntArray[0] = 301

	fmt.Printf("aIntArray=%v, anIntArray=%T\n", anIntArray, anIntArray)

	// declaring an array implicitly
	weekdayTemps := [5]float32{101.4, 98.4, 91.3, 100.4, 95.6}
	fmt.Printf("weekdayTemps=%v, weekdayTemps=%T\n", weekdayTemps, weekdayTemps)

	// arrays are not supported even with sorting functions
	// sort.Float64Slice(weekdayTemps)

	// where is array useful then? may be only where fixed size
	// restrictions are required?
}
