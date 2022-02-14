// Exploring mutex from built-in sync package
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Ledger struct {
	// Keep track of account balances
	balances map[string]float64

	// number of transactions
	nTransactions int

	// To provide safe update to the balances
	mutex sync.Mutex
}

func (l *Ledger) Update(id string, val float64) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.balances[id] += val
	l.nTransactions++
}

func (l *Ledger) String() string {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return fmt.Sprintf("%+v", l.balances)
}

func (l *Ledger) ShowBalance() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for k, v := range l.balances {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Printf("Number of transactions: %v\n", l.nTransactions)
}

type Transaction struct {
	name string
	val  float64
}

func getTransactions(file string) []Transaction {
	// Read the transactions from a file
	buf, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var transactions []Transaction
	for _, line := range strings.Split(string(buf), "\n") {
		fields := strings.Fields(line)
		val, _ := strconv.ParseFloat(fields[1], 64)

		transactions = append(transactions, Transaction{name: fields[0], val: val})
	}
	return transactions
}

func main() {

	transactions := getTransactions("transactions.txt")
	l := Ledger{balances: make(map[string]float64)}

	settle := func(transactions []Transaction) {
		for _, t := range transactions {
			l.Update(t.name, t.val)
		}

	}

	go settle(transactions[:len(transactions)/2])
	go settle(transactions[len(transactions)/2:])

	// Giving some time for the transactions to settle.
	// How to sync main to wait for the go routines to complete?
	// through channel? may be for tomorrow.
	time.Sleep(2 * time.Second)
	l.ShowBalance()

	// TODO for tomorrow
	//	- [ ] Update comments
	//	- [ ] Add channel to sync between main and go functions
	//	- [ ] Re-org code
	//	- [ ] Add comment about concurrent map writes error when running without lock
	//  - [ ] Is there a better way to read line by line? like python file iterator?
}
