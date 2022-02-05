// Learning string operations, manipulation, conversions.
package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1, s2 := "foo", "bar"

	// string comparison with operators
	// They compare the strings lexcially, and return the result
	// as boolean
	fmt.Printf("s1 < s2 : %v\n", s1 < s2)
	fmt.Printf("s1 == s2 : %v\n", s1 == s2)

	// string comparison with strings.Compare
	// Compare works much like strcmp
	// returns 0 if equal, -1 if s1 < s2, 1 if s1 > s2
	// Compare is included only for symmetry with bytes
	// Use the comparison instead.
	fmt.Printf("strings.Compare(s1, s2): %v\n", strings.Compare(s1, s2))

	// string Join
	words := []string{s1, s2}
	s3 := strings.Join(words, "==")
	fmt.Printf("s3: %v\n", s3)

	// string Split
	fmt.Printf("strings.Split(s3, \"==\"): %v\n", strings.Split(s3, "=="))

	line := "the quick brown fox jumps over the lazy dog"

	// string contains
	fmt.Printf("strings.Contains(line, \"the\"): %v\n", strings.Contains(line, "the"))

	// string startswith, endswith
	// Ooops.. I need to stop talking in Python here
	// In Go, it is HasPrefix, HasSuffix
	fmt.Printf("strings.HasPrefix(line, \"kdkd\"): %v\n", strings.HasPrefix(line, "kdkd"))
	fmt.Printf("strings.HasSuffix(line, \"dog\"): %v\n", strings.HasSuffix(line, "dog"))

	// split by spaces
	words = strings.Fields(line)
	fmt.Printf("words: %v\n", words)

	beatIt := `They're out to get you, better leave while you can
	Don't wanna be a boy, you wanna be a man
	You wanna stay alive, better do what you can
	So beat it, just beat it`

	// split by specific fields
	splitByPunc := func(r rune) bool {
		return unicode.IsPunct(r)
	}

	fmt.Printf("strings.FieldsFunc(beatIt, splitByPunc): %v\n",
		strings.FieldsFunc(beatIt, splitByPunc))

	// substring search
	fmt.Printf("strings.Index(beatIt, \"better\"): %v\n",
		strings.Index(beatIt, "better"))

	fmt.Printf("strings.ToTitle(line): %v\n", strings.ToTitle(line))
	fmt.Printf("strings.ToLower(line): %v\n", strings.ToLower(line))

	// so many other functions to explore in strings
	// Replacer, Reader, Trim, etc.

	// Reader is used in many different places. one such example
	// to decode json from strings
	aJsonString := `{"name": "foo"}`
	decoder := json.NewDecoder(strings.NewReader(aJsonString))
	var m map[string]string

	_ = decoder.Decode(&m)
	fmt.Printf("m: %v\n", m)
}
