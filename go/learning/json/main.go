// Learning about json encoding and decoding

package main

import "log"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	UnmarshalJson()
	UnmarshalJsonUnknownKeys()

	// Unmarshal parses the json encoded data and stores the result into the given value
	// Decoder reads the next json encoded value from the given input stream
	// Decoder internally uses unmarshal

	DecodeJsonValue()
	DecodeJsonList()
	DecodeJsonListIterative()
	DecodeJsonFromWeb()
}
