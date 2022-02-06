package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Empty interfaces are interface types with no methods
// they can hold value of any type. In Go, every type implements
// at least zero methods. So any type can be considered an interface.
// This allows the code to use empty interfaces to handle values of
// unknown type.
//
// fmt.Print functions are a good example for the usage of empty interfaces.
// fmt.Println is a variadic function that takes arbitrary number of input
// values of type interface{}
//
// func Println(a ...interface{}) (n int, err error) {
// 	return Fprintln(os.Stdout, a...)
// }
//
func EmptyInterface() {
	fmt.Println("some string")

	// Another example for empty interface
	// if we know the type of value ahead, we can define
	// a map with known type e.g. map[string]string
	// sometimes with json data, we do not know the types
	// ahead or there could be values of different types

	var jsonData map[string]interface{}

	jsonString := `{"city": "Phoenix", "zip": 85001, "density": 3102.92}`
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	err := decoder.Decode(&jsonData)
	if err != nil {
		panic(err)
	}

	for k, v := range jsonData {
		fmt.Printf("key:%v, value:%v, type(value):%T\n", k, v, v)
	}
}
