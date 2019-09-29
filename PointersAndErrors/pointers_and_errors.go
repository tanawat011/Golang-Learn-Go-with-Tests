package PointersAndErrors

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return (*w).balance
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount BitCoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
