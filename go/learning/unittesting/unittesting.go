// Exploring unit testing
// Run tests using "go test"
// testing package provides the tools to run the unit tests
// test files are typically named as file_test.go
// tests are created by creating a function with name
// beginning with Test, that takes *testing.T as input.
// testing.T.Error report test failures but continue to
// run other tests. testing.T.Fatal reports the failure
// and stops the test immediately.

package main

import (
	"math"
	"testing"
)

func MostCommon(text string) string {
	freq := make(map[rune]int)
	for _, c := range text {
		freq[c] += 1
	}

	var mostCommonChar rune
	var maxFreq int

	maxFreq = int(math.Inf(-1))

	for char, f := range freq {
		if f > maxFreq {
			maxFreq = f
			mostCommonChar = char
		}
	}

	return string(mostCommonChar)
}

func TestMostCommon(t *testing.T) {
	// TODO
}
