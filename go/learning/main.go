package main

// TODO
// Add package comments

import (
	"github.com/deepns/codegarage/go/learning/files"
	"github.com/deepns/codegarage/go/learning/http"
)

func main() {
	// TODO
	// Bring in other files under here

	// File operations
	files.BasicReading()
	files.BasicWriting()
	files.ReadUsingIOUtil()

	// exploring http package
	// Make a simple GET request. Read all the data.
	http.GetAllAndDecodeJson()

	// Instead of getting all raw data, converting them into strings
	// and then converting them to strings to run json decoder,
	// resp Body supports the Reader interface, so we can run json
	// decoder directly on the response Body too.
	http.GetAndDecodeJson()
}
