package main

/*
 * Learning about
 * - Reading from standard input using a buffered reader.
 * - String conversions using strconv package
 * - String operations using methods in strings package. E.g. TrimSpace
 */

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Learning inputs...Enter radius, get area")
	fmt.Print("Enter radius:")

	// read from standard input using a buffered reader
	reader := bufio.NewReader(os.Stdin)

	// specify the delimited in bytes. "" is considered a string in Go. Use '' for bytes
	radiusStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert the radius in string format to a float
	// trim any trailing spaces in the input
	radius, err := strconv.ParseFloat(strings.TrimSpace(radiusStr), 64 /* bitsize */)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Specify the floating point precision in the formatting verb
	fmt.Printf("Area of circle with radius %v = %.2f\n",
		radius, math.Pi*math.Pow(radius, 2))
	
	// reading values using scanner
	fmt.Println("Enter some numbers:")
	scanner := bufio.NewScanner(os.Stdin)
	// read until EOF
	// max token size is determined by bufio.MaxScanTokenSize
	for scanner.Scan() {
		fmt.Printf("You entered: %v\n", scanner.Text())
	}
}
