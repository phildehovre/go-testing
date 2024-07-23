package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds for withdrawal")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(q Bitcoin) Bitcoin {
	w.balance = w.balance + q
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(q Bitcoin) (err error) {
	if w.balance < q {
		return ErrInsufficientFunds
	}
	w.balance -= q
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
