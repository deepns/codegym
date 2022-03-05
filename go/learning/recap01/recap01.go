// Revisiting some things that I learned in the past few weeks
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	// Seed the random generator with current unix timestamp
	rand.Seed(time.Now().Unix())

	// Generate a slice of random numbers
	randNumbers := []int{
		rand.Intn(10000),
		rand.Intn(10000),
		rand.Intn(10000),
		rand.Intn(10000),
		rand.Intn(10000),
	}

	// append to a slice
	for i := 0; i < 5; i++ {
		randNumbers = append(randNumbers, rand.Intn(1000))
	}

	// iterate a slice
	for _, n := range randNumbers {
		fmt.Printf("n: %v\n", n)
	}

	// sleep for sometime, time specified in nanosecond durations
	time.Sleep(time.Second / 2) // sleeping for half a second

	// built in time units
	// time.Duration supports Stringer
	// so duration is printed in more readable format than raw integer values
	fmt.Printf("time.Second: %v\n", time.Second)
	fmt.Printf("time.Millisecond: %v\n", time.Millisecond)
	fmt.Printf("time.Microsecond: %v\n", time.Microsecond)

	// nameless struct
	ports := []struct {
		name   string
		number int
	}{
		{"http", 80},
		{"https", 443},
		{"ssh", 22},
	}

	fmt.Printf("ports: %T\n", ports)
	fmt.Printf("ports: %v\n", ports)

	// split strings
	line := "the quick brown fox jumps over lazy dog"
	for _, word := range strings.Split(line, " ") {
		fmt.Printf("word: %v\n", word)
	}

	csv := "apple,pear,orange,banana"
	for i, fruit := range strings.Split(csv, ",") {
		fmt.Printf("fruit[%d]: %v\n", i, fruit)
	}

	// create a new logger
	recapLog := log.New(os.Stdout, "[ recap ] ", log.LstdFlags)
	recapLog.Println("Created the logger")

	// Get current time
	cTime := time.Now()
	fmt.Printf("cTime: %v\n", cTime)

	tomorrow := cTime.AddDate(0, 0, 1)
	fmt.Printf("tomorrow: %v\n", tomorrow)

	// formatting time values
	fmt.Printf("cTime.Format(time.Stamp): %v\n", cTime.Format(time.Stamp))
	fmt.Printf("cTime.Format(time.RFC3339): %v\n", cTime.Format(time.RFC3339))
}
