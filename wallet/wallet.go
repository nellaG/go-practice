package main

import "fmt"

type Wallet struct {
	balance int // private (starts with lowercase)
}

// method
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
