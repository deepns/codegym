package maps

// Learning maps

import (
	"fmt"
	"sort"
)

// Basics of maps
//	defininig maps
//	make maps
//	iterating
//	insertion, deletions
func Basics() {
	fmt.Println("=========== Learning Maps ============")
	// declare the map type explicitly
	// var varName map[keyType]valueType
	var zipcodes map[string]int
	fmt.Printf("type(zipcodes): %T, zipcodes: %v\n", zipcodes, zipcodes)

	// no memory allocated for the map yet. so can't
	// assign a value to a nil map
	// will panic with
	// panic: assignment to entry in nil map
	// zipcodes["San Jose"] = 87098
	zipcodes = make(map[string]int)

	// adding new value into a map and modifiying existing value
	// are similar to how maps are handled in other languages
	zipcodes["San Jose"] = 87098
	fmt.Printf("type(zipcodes): %T, zipcodes: %#v\n", zipcodes, zipcodes)

	// create map with make
	weather := make(map[string]float64)
	weather["monday"] = 34.5
	weather["tuesday"] = 29.4
	// overwrite the value of an existing key
	weather["monday"] = 30.2
	fmt.Printf("type(weather):%T, weather: %v\n", weather, weather)

	// int-to-string map
	// declared implicitly
	fileDescriptors := map[int]string{
		0: "stdin",
		1: "stdout",
		2: "stderr",
	}

	// len() works with maps too
	fmt.Printf("fileDescriptors: %v, len(fileDescriptors): %v\n", fileDescriptors, len(fileDescriptors))

	// iterating maps with range keyword.
	// if only one variable is provided, range gives the key
	// NOTE: maps are unordered. so no order guaranteed when iterating
	// with range keyword
	for day := range weather {
		fmt.Printf("key:%v, value:%v\n", day, weather[day])
	}

	// if two variables are provided, range gives both key and value
	for k, v := range fileDescriptors {
		fmt.Printf("%v => %v\n", k, v)
	}

	// Deleting a key is done using the built in delete function
	delete(zipcodes, "San Jose")
	fmt.Printf("after delete: zipcodes: %v\n", zipcodes)

	// if key is non-existent or nil, then delete is a no-op. no error raised.
	delete(zipcodes, "Acrabadra")

	// get a slice of keys in the map
	keys := make([]int, 0, len(fileDescriptors))
	for k := range fileDescriptors {
		keys = append(keys, k)
	}
	fmt.Printf("keys: %v, sorted(keys):%v\n", keys, sort.IntsAreSorted(keys))

	// clearing a map
	for k := range fileDescriptors {
		delete(fileDescriptors, k)
	}
}
