package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// DecodeJsonValue decodes a stream of json data into structs
func DecodeJsonValue() {

	fmt.Println("========== Decoding stream of json values ==========")

	// Going to decode this json string into structs
	const jsonStream = `
	{"name": "Mix", "frequency": 93.4, "zip": 74749}
	{"name": "Rock", "frequency": 101.4, "zip": 73339}
	{"name": "Jazz", "frequency": 91.4, "zip": 53563}
	{"name": "News", "frequency": 91.5, "zip": 38474}`

	fmt.Printf("Decoding jsonStream: %v\n", jsonStream)

	// Define the message struct
	// this struct can have all or partial set of
	// keys of the json data we want to decode
	// At first, this seemed like a pain to require the
	// definition of schema of the json being decoded.
	// Unlike Python, decoding json from a stream doesn't
	// give a dict here. Instead, we get a struct. In a way
	// it is good though. It makes the programmer intent explicit
	// about what needs to be decoded.

	// Decoder uses json.Unmarshal underneath. Snippet from Unmarshal docs
	// To unmarshal JSON into a struct, Unmarshal matches incoming object keys
	// to the keys used by Marshal (either the struct field name or its tag),
	// preferring an exact match but also accepting a case-insensitive match.
	// By default, object keys which don't have a corresponding struct field are
	// ignored (see Decoder.DisallowUnknownFields for an alternative).

	type Station struct {
		Name      string
		Frequency float64
		Zip       int
	}

	// json decoder works on top of types that acts a Reader
	// so make a reader for the string using strings.NewReader
	decoder := json.NewDecoder(strings.NewReader(jsonStream))

	// looping until the reader reaches EOF
	for {
		var station Station
		if err := decoder.Decode(&station); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("station: %+v\n", station)
	}
}

// DecodeJsonList decodes a json list into a go slice
func DecodeJsonList() {
	fmt.Println("========== Decoding a json list into a slice ==========")

	const jsonList = `
	[{"name": "Mix", "frequency": 93.4, "zip": 74749},
	{"name": "Rock", "frequency": 101.4, "zip": 73339},
	{"name": "Jazz", "frequency": 91.4, "zip": 53563},
	{"name": "News", "frequency": 91.5, "zip": 38474}]`

	type Station struct {
		Name      string
		Frequency float64
	}

	// Decoding the full data in a slice
	var stations []Station
	decoder := json.NewDecoder(strings.NewReader(jsonList))
	if err := decoder.Decode(&stations); err != nil {
		panic(err)
	}

	for _, station := range stations {
		fmt.Printf("station: %v\n", station)
	}

}

// DecodeJsonListIterative decodes values from json list iteratively
// instead of reading all at once into a slice.
func DecodeJsonListIterative() {
	fmt.Println("========== Decoding a json list iteratively ==========")

	// Decoding directly into a slice can be expensive
	// on larger data. Instead, iterate over the json
	// list
	const jsonList = `
	[{"name": "Mix", "frequency": 93.4, "zip": 74749},
	{"name": "Rock", "frequency": 101.4, "zip": 73339},
	{"name": "Jazz", "frequency": 91.4, "zip": 53563},
	{"name": "News", "frequency": 91.5, "zip": 38474}]`

	type Station struct {
		Name      string
		Frequency float64
	}

	decoder := json.NewDecoder(strings.NewReader(jsonList))

	// read the open bracket [
	token, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	// Note the type of token. it is not string
	// it is json.Delim
	fmt.Printf("%T: %v\n", token, token)

	// read until there is more data under this delimiter
	for decoder.More() {
		var station Station
		if err := decoder.Decode(&station); err != nil {
			panic(err)
		}
		fmt.Printf("station: %+v\n", station)
	}

	// read the closing bracket ]
	token, err = decoder.Token()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T: %v\n", token, token)
}

func DecodeJsonFromWeb() {
	fmt.Println("========== Decoding a json data pulled from web ==========")

	const httpbin = "https://httpbin.org/get"

	resp, err := http.Get(httpbin)
	checkErr(err)
	defer resp.Body.Close()

	var data map[string]interface{}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&data)
	checkErr(err)
	for k, v := range data {
		fmt.Println(k, ":", v)
	}
}
