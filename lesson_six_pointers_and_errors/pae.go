package pae

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("can not withdraw more money than the current balance")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(ammount Bitcoin) {
	w.balance += ammount
}

func (w *Wallet) Withdraw(ammount Bitcoin) error {
	if w.balance < ammount {
		return ErrInsufficientFunds
	}
	w.balance -= ammount
	return nil
}
