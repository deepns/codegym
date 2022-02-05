// Exploring go functions
// closures, function variables, anonymous functions
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

// GetIdGen returns a sequence number generator starting
// at the given start. This is an example of using closures.
func GetIdGen(startAt int) func() int {
	start := startAt
	return func() int {
		start += 1
		return start
	}
}

func main() {

	// functions can be treated like any other types
	// can be even assigned to variables
	var greet = func(msg string) {
		fmt.Println("Bonjour, ", msg)
	}
	greet("Alice")

	nextId := GetIdGen(rand.Intn(100))
	for i := 0; i < 10; i++ {
		fmt.Printf("nextId(): %v\n", nextId())
	}

	nextId2 := GetIdGen(rand.Intn(100))
	for i := 0; i < 10; i++ {
		fmt.Printf("nextId2(): %v\n", nextId2())
	}

	// Two functions having the same argument types and return types
	// are considered identical

	// Passing around functions
	// splitting strings by a specific character is one good example
	// strings.FieldsFunc takes a function of type func(r rune) bool
	wordsWithPunc := "the, quick, brown fox stole the dog's food"
	filter := func(r rune) bool {
		return unicode.IsPunct(r)
	}
	fmt.Printf("filter: %T\n", filter)
	fmt.Printf("strings.FieldsFunc(wordsWithPunc, filter): %v\n",
		strings.FieldsFunc(wordsWithPunc, filter))

	// functions can be created on the fly like this too
	strings.FieldsFunc(wordsWithPunc, func(r rune) bool { return unicode.IsSpace(r) })

	// Anonymous functions
	// functions can be created and called at the same time
	wordCount := func(words string) int {
		return len(strings.Fields(words))
	}(wordsWithPunc)

	fmt.Printf("wordCount: %v\n", wordCount)
}
