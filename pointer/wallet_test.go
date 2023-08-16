package pointer

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("expect %d but got %d", want, got)
	}

}
