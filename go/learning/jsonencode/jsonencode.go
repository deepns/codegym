// Exploring json encoding
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Response struct {
	NumRecords int      `json:"num_records"`
	Records    []string `json:"records"`
	Version    float32
}

// EncodeJsonStructs talks about encoding structs into json format
// Only public fields of struct are marshaled into json data
// Json tags can be specified optionally to use a different key instead
// of field name
func EncodeJsonStructs() {
	fmt.Println("\n========== EncodeJsonStructs ============")

	faang := Response{5, []string{
		"facebook",
		"apple",
		"amazon",
		"netflix",
		"google",
	},
		1.5}

	// Marshal by default doesn't include any indents or prefix to keep the
	// data compact. If indenting is needed, use MarshalIndent
	b, err := json.Marshal(faang)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b): %v\n", string(b))

	// Marshaling with Indent
	b, _ = json.MarshalIndent(faang, "", "  ")
	fmt.Printf("with indenting: string(b):\n%v\n", string(b))
}

// EncodeJsonBasic talks about encoding go basic types in json format
func EncodeJsonBasic() {
	fmt.Println("\n========== EncodeJsonBasic ============")
	// go has built in encoder for all basic types.
	// json.Marshal takes the argument, encodes into json formatted byte array
	b, err := json.Marshal("hello")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("b: %v\n", b)
	fmt.Printf("string(b): %v\n", string(b))

	// marshaling ints and floats
	b, _ = json.Marshal(10)
	fmt.Printf("string(b): %v\n", string(b))
	b, _ = json.Marshal(13.45)
	fmt.Printf("string(b): %v\n", string(b))
	// marshaling boolean
	b, _ = json.Marshal(true)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("string(b): %v\n", string(b))
}

// EncodeJsonToStream talks about encoding types to an output stream
// Uses json.Encoder instead of Marshal
func EncodeJsonToStream() {
	fmt.Println("\n========== EncodeJsonToStream ============")

	var sb strings.Builder

	// NewEncoder takes any writer. so even os.Stdout works good here.
	encoder := json.NewEncoder(&sb)
	// indent by 2 spaces
	encoder.SetIndent("", "  ")

	// encoding different values into json format
	i := 10
	err := encoder.Encode(i)
	if err != nil {
		log.Fatal(err)
	}

	// Encoding rarely runs into error, so ignoring the error handling for now
	encoder.Encode("hello")
	encoder.Encode(map[string]interface{}{
		"status": "Ok",
		"code":   200,
		"error":  false,
		"result": map[string]int{
			"x": 10,
			"y": 20,
		},
		"versions": []float64{
			12.74,
			12.73,
			12.71,
		},
	})
	encoder.Encode([]float64{
		10.12,
		13.44,
		424.2424,
		2424.222,
	})
	fmt.Println(sb.String())
}

func main() {
	EncodeJsonBasic()
	EncodeJsonStructs()
	EncodeJsonToStream()
}
