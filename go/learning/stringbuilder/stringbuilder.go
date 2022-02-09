// Exploring string builder from strings package
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	// declaring the string builder
	//	string builder has a self referencing pointer and a buffer.
	//	both are unexported.
	var sb strings.Builder

	fmt.Printf("sb: %v\n", sb)
	fmt.Printf("sb.Cap(): %v\n", sb.Cap())
	fmt.Printf("sb.Len(): %v\n", sb.Len())

	// Writing some strings into the builder
	// WriteString returns number of bytes written, and a nil error
	// ignoring the return values for now
	sb.WriteString("hello,")
	sb.WriteString("string builder!")

	// Get the built string with String()
	fmt.Printf("sb.String(): %v\n", sb.String())

	// check the length and capacity after adding some strings
	fmt.Printf("sb.Len(): %v\n", sb.Len())
	fmt.Printf("sb.Cap(): %v\n", sb.Cap())

	// If we know how much capacity will be needed ahead, we
	// can grow the builder
	// for e.g growing by 512 bytes
	sb.Grow(512)

	// check the length and capacity after growing
	fmt.Printf("sb.Len(): %v\n", sb.Len())
	fmt.Printf("sb.Cap(): %v\n", sb.Cap())

	// Writing some raw bytes into the builder
	sb.Write([]byte("...\tsome byte array"))

	// clear the buffer
	sb.Reset()
	fmt.Printf("sb:%v, len:%v, cap:%v\n", sb.String(), sb.Len(), sb.Cap())

	// Printing some other types
	mixedTypes := make(map[string]interface{})

	mixedTypes["int"] = 20
	mixedTypes["float"] = 10.34
	mixedTypes["bool"] = false

	for k, v := range mixedTypes {
		fmt.Fprintf(&sb, "k: %v, v:%v\n", k, v)
	}

	fmt.Printf("sb.String(): %v\n", sb.String())

	// I first thought this would be equivalent to sb.String() since
	// sb is a Stringer and the String() method would be called
	// during fmt.Print. Looking into the method signature cleared
	// the doubt.

	// Though StringBuilder has a String() defined, the receiver is
	// a pointer
	// func (b *Builder) String() string {
	//	return *(*string)(unsafe.Pointer(&b.buf))
	// }

	// So this would print the go representation of the StringBuilder type
	fmt.Printf("sb.String(): %v\n", sb)

	// whereas, this would print the string value of the StringBuilder
	// equivalent to printing the return value of sb.String()
	fmt.Printf("sb.String(): %v\n", &sb)

	// Since sb is a Writer (it has Write() method), this can also
	// be used in place where io.Writer is used
	var w io.Writer = &sb
	w.Write([]byte("from ioWriter"))
	fmt.Printf("sb.String(): %v\n", sb.String())

	sb.Reset()

	// Fprintf writes a formatted string into the io Writer
	// Note that we pass the address of the StringBuilder in place
	// of Writer, because the receiver of Write() is a pointer
	fmt.Fprint(&sb, "formatted print...")
	fmt.Fprintf(&sb, "len: %v, cap:%v", sb.Len(), sb.Cap())
	fmt.Println(sb.String())

	// Builder can also be used with pointers
	sb2 := new(strings.Builder)
	sb2.WriteString("Another string builder\n")
	fmt.Fprintf(sb2, "len: %v", sb.Len())
	fmt.Printf("sb2.String(): %v\n", sb2.String())

	// some examples of string builder usage
	// https://github.com/google/go-containerregistry/blob/62eaac05655962b22cd35ec041350f69941f4d47/pkg/v1/remote/transport/ping.go#L152
	// https://github.com/google/go-github/blob/b5776f9fd95335bec17c498b5f259b4fd30af589/update-urls/main.go#L640
}
