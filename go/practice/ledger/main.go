// This is a sample exercise to put some Go concepts in use.
// The file transactions.txt has thousands of transactions, one per line.
// The transactions are between three people and the net balance is 0.
// This exercise builds a ledger to settle the transactions.
// The following Go concepts are used in solving this exercise:
//	- structs
//	- maps - creation, iteration, insertion
//	- slices - creation, iteration, append
//	- struct methods
//	- function types
//	- function with named return values
//	- constants
//	- type conversions
//	- errors - creating errors with fmt.Errorf
//	- mutex - to protect the map from concurrent writes
//	- get current time in milli, micro and nanoseconds using time package
//	- go routines - to update the ledger concurrently
//	- channels - to communicate between main and the go routines about their completion
//	- defer
//	- read a file line by line using bufio scanner
//	- reading environment variables using os package
//	- string conversion using strconv package
//	- split strings using strings package
//  - generate docs with 'go doc'
//	- command line flags
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// getTransactions reads the transactions from the given file and returns them
// in a slice
func getTransactions(file string) []Transaction {
	var transactions []Transaction

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Use a scanner to read the file line by line
	// scanner uses bufio.ScanLines as the default split function
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		val, _ := strconv.ParseFloat(fields[1], 64)
		transactions = append(transactions, Transaction{name: fields[0], val: val})
	}

	// it would be even better if we split the file directly between the workers
	// instead of building records for the entire file in memory and then splitting
	// it. That would require careful data alignment though.

	return transactions
}

// getWorkerCount returns the number of worker routines
func getWorkerCount() (int, error) {
	nWorkers := 2
	const MAX_WORKERS = 16

	// just for fun, get the number of workers through an environment variable
	_, ok := os.LookupEnv("LEDGER_WORKERS")
	if ok {
		n, err := strconv.ParseInt(os.Getenv("LEDGER_WORKERS"), 10, 0)
		if err != nil {
			return nWorkers, err
		}
		if n <= 0 || n > MAX_WORKERS {
			return nWorkers, fmt.Errorf("invalid number of workers: %v", n)
		}
		nWorkers = int(n)
	}
	return nWorkers, nil
}

func NewLedger() *Ledger {
	return &Ledger{db: make(map[string]float64)}
}

func main() {
	transFile := flag.String("file", "transactions.txt", "source file to read the transactions")
	flag.Parse()

	transactions := getTransactions(*transFile)
	nWorkers, err := getWorkerCount()
	if err != nil {
		log.Fatal(err)
	}

	l := NewLedger()
	l.Settle(transactions, nWorkers)
	l.ShowBalance()
}
