// Exploring sorting functions
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Point struct {
	x, y int
}

func main() {

	rand.Seed(time.Now().Unix())

	// Generating some random numbers
	ints := []int{
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
	}

	words := []string{
		"abcde",
		"hello",
		"world",
		"foo",
		"xoxo",
	}

	// sort them using functions from built in sort package
	// sort package has predefined functions to sort int, float64 and string slices in ascending order
	// sort.Ints() sorts the slice in increasing order
	sort.Ints(ints)
	sort.Strings(words)

	fmt.Printf("ints: %v\n", ints)
	fmt.Printf("words: %v\n", words)

	// Check a slice is sorted are not
	words = append(words, "onemore")
	fmt.Printf("words: %v\n", words)
	fmt.Printf("sort.StringsAreSorted(words): %v\n", sort.StringsAreSorted(words))

	// sort in reverse order using sort.Reverse
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Printf("ints: %v\n", ints)

	// another way to sort in reverse by providing a custom less function
	sort.Slice(words, func(i, j int) bool {
		return words[i] > words[j]
	})
	fmt.Printf("words: %v\n", words)

	sort.Ints(ints)
	fmt.Printf("ints: %v\n", ints)

	// sort also provides functions to search using binary search on a slice
	// sorted in ascending order. Search returns the index to insert the search value
	fmt.Printf("sort.SearchInts(ints, 25): %v\n", sort.SearchInts(ints, 25))

	points := []Point {
		{rand.Intn(100), rand.Intn(100)},
		{rand.Intn(100), rand.Intn(100)},
		{rand.Intn(100), rand.Intn(100)},
		{rand.Intn(100), rand.Intn(100)},
		{rand.Intn(100), rand.Intn(100)},
	}

	// for slices of any type, sort.Slice with a custom less function works great
	// sorting the points slice by the x coordinates.
	sort.Slice(points, func(i, j int) bool { return points[i].x < points[j].x})
	fmt.Printf("points: %v\n", points)
}
