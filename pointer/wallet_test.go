package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertionBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("test deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertionBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test withdraw bit coin", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertionBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test insufficient fund", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Error("expect error")
		}
	})

}
