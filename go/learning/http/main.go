package main

func main() {
	// exploring http package
	// Make a simple GET request. Read all the data.
	GetAllAndDecodeJson()

	// Instead of getting all raw data, converting them into strings
	// and then converting them to strings to run json decoder,
	// resp Body supports the Reader interface, so we can run json
	// decoder directly on the response Body too.
	GetAndDecodeJson()

	UrlParsing()

	PostSomething()
}
