package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		
		if got != want {
			t.Errorf("\ngot %v\nwant %v", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string){
		t.Helper()
		if got == nil {
			t.Fatal("didn't get error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q\nwant %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{
			balance: Bitcoin(20),
		}

		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdram insuficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(10))

		assertError(t, err, "cannot withdraw, insuficient funds")
		assertBalance(t, wallet, startingBalance)
	})
}
