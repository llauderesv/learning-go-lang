package main

import (
	"fmt"
	"sync"
)

type Bank struct {
	balance int
}

func (b *Bank) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(b.balance)
	b.balance = b.balance + amount
	fmt.Println("Current Balance: ", b.Balance())
}

func (b *Bank) Balance() int {
	return b.balance
}

func main() {

	b := Bank{balance: 100}
	var wg sync.WaitGroup

	wg.Add(1)
	go b.Deposit(200, &wg)
	wg.Wait()

	wg.Add(1)
	go b.Deposit(100, &wg)
	wg.Wait()
}
