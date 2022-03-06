// Recap - part 02
//	init functions
//	custom errors
//	reading inputs with
//		Scanf
//		bufio.NewReader
//		bufio.NewScanner
// string-to-int conversion using strconv.ParseInt
// trimming strings with strings.TrimSpace

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var initCount int

// init() runs after completing all the imports (and the init calls within the
// imports) and before running main
func init() {
	initCount = 1
}

// init can be defined multiple times and is called in the order in which
// it is defined
func init() {
	initCount = 2
}

func init() {
	initCount = 3
}

func init() {
	initCount = 4
}

// Defining custom errors

type NegativeNumberError int

func (n NegativeNumberError) Error() string {
	return fmt.Sprintf("%v is negative. wanted a postive integer", int(n))
}

func main() {
	fmt.Printf("initCount: %v\n", initCount)
	
	// Read input from stdin
	var aNum int64
	fmt.Println("Enter some positive integer: ")
	fmt.Scanf("%d", &aNum)

	if aNum <= 0 {
		fmt.Println(NegativeNumberError(aNum))
	}

	// another way to read inputs...through bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter another number:")
	aNumStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	// convert string to a byte array
	fmt.Printf("bytes(aNumStr): %v\n", []byte(aNumStr))

	// convert number in string format to integer
	// trim the new line and space
	aNum, err = strconv.ParseInt(strings.TrimSpace(aNumStr), 10, 64)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("aNum: %v\n", aNum)

	fmt.Println("Enter some numbers. one per line")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("you entered: %v\n", scanner.Text())
	}
}
