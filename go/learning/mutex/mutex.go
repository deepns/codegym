// Exploring mutex from built-in sync package
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Transaction struct {
	name string
	val  float64
}

// Ledger keep tracks of the transactions
type Ledger struct {
	// Keep track of account db
	db map[string]float64

	// number of transactions
	nTransactions int

	// To provide safe update to the balances
	mutex sync.Mutex
}

// Update updates the ledger with the given transaction
func (l *Ledger) Update(t Transaction) {
	// If the map is updated without the lock by multiple go routines,
	// it often runs into "fatal error: concurrent map writes" error
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.db[t.name] += t.val
	l.nTransactions++
}

// String returns the string representation of the ledger
func (l *Ledger) String() string {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return fmt.Sprintf("db: %+v, nTransactions:%v", l.db, l.nTransactions)
}

// ShowBalance prints the current ledger balance to stdout
func (l *Ledger) ShowBalance() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for k, v := range l.db {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Printf("Number of transactions: %v\n", l.nTransactions)
}

// getTransactions returns a slice of transactions read from the given file
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
func getWorkerCount() int {
	nWorkers := 1
	const MAX_WORKERS = 16

	// just for fun, get the number of workers through an environment variable
	_, ok := os.LookupEnv("LEDGER_WORKERS")
	if ok {
		n, err := strconv.ParseInt(os.Getenv("LEDGER_WORKERS"), 10, 0)
		if err == nil {
			if n > 0 {
				nWorkers = int(math.Min(float64(n), MAX_WORKERS))
			}
		}
	}
	return nWorkers
}

func main() {

	transactions := getTransactions("transactions.txt")
	l := Ledger{db: make(map[string]float64)}
	c := make(chan int)
	settle := func(transactions []Transaction, workerId int, c chan int) {
		for _, t := range transactions {
			l.Update(t)
		}
		// Send the workerid over the channel when done updating.
		c <- workerId
	}

	nWorkers := getWorkerCount()
	if len(transactions)%nWorkers > 0 {
		fmt.Println("work is not evenly divisible. Adding an extra worker!")
		nWorkers++
	}
	fmt.Println("Splitting the work between", nWorkers, "workers")

	// launch a go routine for each chunk of data
	chunkSize := len(transactions) / nWorkers
	for i := 0; i < nWorkers; i++ {
		offset := i * chunkSize
		size := offset + chunkSize

		// Make sure we don't go out of bounds
		if offset+chunkSize >= len(transactions) {
			size = len(transactions)
		}

		go settle(transactions[offset:size], i, c)
	}

	// Show the balance at the end
	defer l.ShowBalance()

	// wait until all transactions are settled
	for i := 0; i < nWorkers; i++ {
		// can see workers completing different order during each invocation
		worker := <-c
		fmt.Printf("worker%v done!\n", worker)
	}
}
