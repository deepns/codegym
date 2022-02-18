package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Transaction struct {
	name string
	val  float64
}

// Ledger keeps a record of the transactions
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
	fmt.Printf("Net balance: %v\n", l.TotalBalance())
}

// Settle settles the transactions by updating the ledger using given number of workers
func (l *Ledger) Settle(transactions []Transaction, nWorkers int) {
	startTime := time.Now().UnixMicro()
	c := make(chan int)
	settle := func(transactions []Transaction, workerId int, c chan int) {
		for _, t := range transactions {
			l.Update(t)
		}
		// Send the workerid over the channel when done updating.
		c <- workerId
	}

	if len(transactions)%nWorkers > 0 {
		log.Println("work is not evenly divisible. adding an extra worker!")
		nWorkers++
	}

	log.Println("Splitting the work between", nWorkers, "workers")

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

	// wait until all transactions are settled
	for i := 0; i < nWorkers; i++ {
		// can see workers completing different order during each invocation
		worker := <-c
		log.Printf("worker%v done!\n", worker)
	}
	log.Printf("Settled in: %v microsecs\n", (time.Now().UnixMicro() - startTime))
}

// TotalBalance returns the outstanding balance in the ledger
func (l *Ledger) TotalBalance() (totalBalance float64) {
	for _, val := range l.db {
		totalBalance += val
	}
	return
}
