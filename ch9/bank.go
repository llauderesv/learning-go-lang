package main

import "fmt"

type Bank struct {
	balance int
}

func (b *Bank) Deposit(amount int) {
	b.balance = b.balance + amount
}

func (b *Bank) Balance() int {
	return b.balance
}

func main() {

	b := Bank{balance: 100}

	go func() {
		b.Deposit(200)
		fmt.Println("=", b.Balance())
	}()

	go b.Deposit(100)
}
