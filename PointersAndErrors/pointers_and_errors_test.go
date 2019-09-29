package PointersAndErrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{BitCoin(20)}
		wallet.Withdraw(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{BitCoin(20)}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(t, wallet, BitCoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
