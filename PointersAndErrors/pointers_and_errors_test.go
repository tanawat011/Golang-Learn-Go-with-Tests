package PointersAndErrors

import (
	"errors"
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
		err := wallet.Withdraw(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{BitCoin(20)}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(t, wallet, BitCoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})

	t.Run("Withdraw the whole balance", func(t *testing.T) {
		wallet := Wallet{BitCoin(20)}
		err := wallet.Withdraw(BitCoin(20))

		assertNoError(t, err)
		assertBalance(t, wallet, BitCoin(0))
	})
}

func TestBitCoinString(t *testing.T) {
	if got, want := BitCoin(10).String(), "10 BTC"; got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertBalance(t *testing.T, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but didn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
