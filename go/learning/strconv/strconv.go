// Exploring string conversions
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Converting strings to int,float,bool types and vice
	// versa

	// In python str(100) will return "100". Go type conversion doesn't
	// work that way. string(100) will treat 100 as the ascii value instead.
	d := string(100)
	fmt.Printf("d: %v\n", d)

	// int, float, bool in string format can be parsed into their
	// equivalent types with strconv.Parse... functions
	// bitSize 0 corresponds to int
	fmt.Println(strconv.ParseInt("12843", 10 /*base*/, 0 /*bitSize*/))

	// Parsing an invalid integer
	_, err := strconv.ParseBool("1234")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// Parsing a float
	fmt.Println(strconv.ParseFloat("123.23", 64))

	// Atoi is equivalent to ParseInt(str, 10, 0)
	// always convert to 10 base.
	fmt.Println(strconv.Atoi("737367474"))

	// int to strings

	fmt.Printf("strconv.FormatBool(true): %v\n", strconv.FormatBool(true))
	fmt.Printf("strconv.FormatInt(12348, 16): %v\n", strconv.FormatInt(12348, 16))

	// Float to string
	fmt.Printf("strconv.FormatFloat(3884.323, 'f', 2, 64): %v\n",
		strconv.FormatFloat(3884.323, 'f', 2, 64))

	// lot more functions in strconv package to explore
	fmt.Printf("strconv.Quote(\"gopher\"): %v\n", strconv.Quote("gopher"))
}
