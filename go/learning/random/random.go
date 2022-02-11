// Exploring the math/rand package
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Default seed is 1, so the random generator will always
	// generate the same sequence if not seeded differently.
	// Seeding by the system nanosecond
	rand.Seed(time.Now().UnixNano())

	// random integer, between 0 to 2^63-1
	fmt.Printf("rand.Int(): %v\n", rand.Int())

	// like Intn(), there is also Int31n and Int63n
	// versions.

	// random integer, between 0 to n
	fmt.Printf("rand.Intn(100): %v\n", rand.Intn(100))

	// random bytes
	b := make([]byte, 8)
	rand.Read(b)
	fmt.Printf("b: %x\n", b)

	// random float. float is always between 0.0 - 0.1
	fmt.Printf("rand.Float64(): %v\n", rand.Float64())

	// rand.Perm returns a slice of ints with a values
	// between 0..n permutated.
	fmt.Printf("rand.Perm(len(b)): %v\n", rand.Perm(len(b)))

	// Shuffle a slice/array using rand.Shuffle
	headerFields := strings.Fields("NAME       STATUS   ROLES                  AGE   VERSION")
	rand.Shuffle(len(headerFields), func(i, j int) {
		headerFields[i], headerFields[j] = headerFields[j], headerFields[i]
	})
	fmt.Printf("headerFields: %v\n", headerFields)
}
