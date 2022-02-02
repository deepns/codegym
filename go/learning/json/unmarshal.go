package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// UnmarshalJson shows some examples of unmarshaling json encoded
// byte data into defined types
func UnmarshalJson() {
	fmt.Println("========== Unmarshal json data into defined types ==========")

	// json.Unmarshal parses the JSON-encoded data and stores
	// the result in the value pointed to by v (interface type)

	// Decoding a simple json
	data := []byte(`{"day":"monday"}`)
	fmt.Printf("data: %s\n", data)

	valid := json.Valid(data)
	fmt.Printf("valid: %v\n", valid)

	type Date struct {
		Day string
	}

	var d Date
	err := json.Unmarshal(data, &d)
	checkErr(err)
	fmt.Printf("d: %+v\n", d)

	// Another json structured data.
	data = []byte(`{"op": "add", "operands":[10, 20]}`)
	valid = json.Valid(data)
	fmt.Printf("valid: %v\n", valid)

	type Calc struct {
		Op       string
		Operands []int
	}
	var c Calc

	err = json.Unmarshal(data, &c)
	checkErr(err)
	fmt.Printf("c: %+v\n", c)

	// Unmarshaling a json list

	data = []byte(`
		[
			{"op": "add", "operands":[10, 20]}, 
		 	{"op": "sub", "operands":[10, 20]}
		]`)

	var calculations []Calc
	err = json.Unmarshal(data, &calculations)
	checkErr(err)
	fmt.Printf("calculations: %+v\n", calculations)
}

// UnmarshalJsonUnknownKeys shows some example of unmarshaling
// json data into a map. When the keys in the json data are not
// known ahead, we can unmarshal this way to get a map that
// is equivalent to a json.
func UnmarshalJsonUnknownKeys() {
	fmt.Println("========== Unmarshal json data into map ==========")

	const httpbin = "https://httpbin.org/get"

	resp, err := http.Get(httpbin)
	checkErr(err)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	checkErr(err)
	fmt.Printf("data: %s\n", data)

	// Don't know the type ahead. so unmarshall the data
	// into a map
	var dataJson map[string]interface{}
	err = json.Unmarshal(data, &dataJson)
	checkErr(err)
	for k, v := range dataJson {
		fmt.Printf("%s : %v\n", k, v)
	}
}
