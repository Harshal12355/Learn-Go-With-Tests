package main

import (
	"errors"
	"fmt"
)

type Bitcoin int // we create a new type called Bitcoin which is an alias for int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct to hold the balance of Bitcoin
type Wallet struct {
	balance Bitcoin // balance is of type Bitcoin
}

// we add pointers over here with * next to the struct
// when we call this method elsewhere, we are basically copying it again in memory, 
// so when we to "Deposit" money or add the deposit value here to the balance, it 
// is not adding the same balance but instead to a copy of that called in that instance
// To solve this problem, we use pointers to point back to the original balance variable
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount 
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds	// if the amount to withdraw is greater than the balance, return an error
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	wallet := Wallet{} // create a new wallet instance
	wallet.Deposit(10) // deposit 10 BTC

	// Printing the balance
	fmt.Println("Balance:", wallet.Balance()) // should print 10 BTC


	// Withdraw some amount
	err := wallet.Withdraw(5) // withdraw 5 BTC
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Balance after withdrawal:", wallet.Balance()) // should print 5 BTC
	}

	// Try to withdraw more than the balance
	err = wallet.Withdraw(10) // try to withdraw more than the balance
	if err != nil {
		fmt.Println("Error:", err) // should print "insufficient funds"
	}
}