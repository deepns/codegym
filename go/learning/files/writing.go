package main

import (
	"fmt"
	"os"
	"path"
)

// checkErr triggers a panic if err is not nil
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// BasicWriting is about getting started with writing files in go.
// Shows different ways in which a file can be written.
func BasicWriting() {
	fmt.Println("========= Writing files ===========")
	testFile := GetTestFilePath()

	fmt.Println("==> Writing using os.WriteFile")
	data := []byte("S U P E R S T A R")
	// if file exists, WriteFile truncates the file otherwise
	// creates a new file
	err := os.WriteFile(testFile, data, 0755 /*file mode*/)
	checkErr(err)

	fmt.Println("==> Writing using file type, opened using os.Create")
	file, err := os.Create(testFile)
	checkErr(err)
	defer file.Close()
	file.WriteString("G O P H E R S")

	fmt.Printf("Reading from %v which we just wrote\n", path.Base(testFile))
	dataBytes, err := os.ReadFile(testFile)
	checkErr(err)
	fmt.Printf("data: %v\n", string(dataBytes))
}
