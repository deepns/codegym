// This is based on the sample exercise from https://go.dev/tour/methods/23
// A Reader type wrapping another reader is a common pattern in Go.
// Read from a stream of bytes and encrypt them using rot13 cipher
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// rot13 performs rot13 substitution on the given byte of alphabetic character
func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return ((b - 'A' + 13) % 26) + 'A'
	} else if b >= 'a' && b < 'z' {
		return ((b - 'a' + 13) % 26) + 'a'
	} else {
		return b
	}
}

// rot13Reader wraps a io.Reader, and also implements the Reader interface
// When reading, it reads from the underlying reader and modifies the
// the stream by applying rot13 substitution cipher
type rot13Reader struct {
	r io.Reader
}

// Read implements io.Reader, reading plain data from the underlying
// reader and modifies the data using rot13 substituion cipher
func (r *rot13Reader) Read(b []byte) (int, error) {
	// Read from the underlying reader
	n, err := r.r.Read(b)
	if err != nil && err != io.EOF {
		return 0, err
	}

	// Encrypt the data read with Rot13 cipher
	for i := 0; i < len(b); i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("the quick brown fox jumps over lazy dog")
	r := rot13Reader{s}

	// Copy from the reader to the writer (stdout)
	io.Copy(os.Stdout, &r)
	fmt.Println()

	// Reading manually with a byte buffer
	s = strings.NewReader("another string to read")
	r = rot13Reader{s}

	// Reading 8 bytes at a time
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Print(string(b[:n]))
	}

	fmt.Println()
}
