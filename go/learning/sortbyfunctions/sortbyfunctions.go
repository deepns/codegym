// Exploring sorting with custom functions
// functions in sort package sorts the slices by defined order (e.g. ints in
// increasing order, strings in lexicographic order)
// when data needs to be ordered in a different way, the custom logic is
// be implemented through the sort.Interface interface.
// Create a custom type (if sorting the built in types) and implement the
// functions of sort.Interface

package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"sort"
	"time"
)

type byExtension []string

func (f byExtension) Len() int {
	return len(f)
}

func (f byExtension) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f byExtension) Less(i, j int) bool {
	return filepath.Ext(f[i][1:]) < filepath.Ext(f[j][1:])
}

func main() {

	files := []string{
		"a1.txt",
		"foobar.txt",
		"httpserver.go",
		"client.c",
		"onemore.txt",
		"server.c",
		"names.json",
	}

	sort.Sort(byExtension(files))
	fmt.Printf("files (regular sort): %v\n", files)

	// lets shuffle the slice a bit
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(files), func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})

	sort.Stable(byExtension(files))
	fmt.Printf("files (stable sorted): %v\n", files)

	// Here is an example of custom types implemeting the sort.Interface
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.7:src/net/http/header.go;l=149;bpv=1;bpt=1

	// A headerSorter implements sort.Interface by sorting a []keyValues
	// by key. It's used as a pointer, so it can fit in a sort.Interface
	// interface value without allocation.
	// type headerSorter struct {
	// 	kvs []keyValues
	// }

	// func (s *headerSorter) Len() int           { return len(s.kvs) }
	// func (s *headerSorter) Swap(i, j int)      { s.kvs[i], s.kvs[j] = s.kvs[j], s.kvs[i] }
	// func (s *headerSorter) Less(i, j int) bool { return s.kvs[i].key < s.kvs[j].key }

	// another example. from net/rpc
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.7:src/net/rpc/debug.go;bpv=1;bpt=0
}
