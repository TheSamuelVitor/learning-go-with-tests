package pointersanderros

import (
	"testing"
)

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("\ngot %v\nwant %v", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q\nwant %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func TestWallet(t *testing.T) {

	t.Run("test Bitcoin wallet", func(t *testing.T) {
		got := Bitcoin.String(20)
		want := "20 BTC"
		if got != want {
			t.Errorf("got %q\nwant %q", got, want)
		}
	})

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
