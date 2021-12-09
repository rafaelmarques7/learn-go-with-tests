package pae

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		depositAmmount := Bitcoin(10)
		wallet.Deposit(depositAmmount)

		want := depositAmmount

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(4))

		want := Bitcoin(6)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw throws an error when there are unsufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		err := wallet.Withdraw(20)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})
}

func assertError(t testing.TB, errGotten, errExpected error) {
	t.Helper()
	if errGotten == nil {
		// Fatal means the code will exit here
		// Without this, we would do a method call on a nil pointer, in the next if block, and the code would panic
		t.Fatal("Expected an error but did not get one")
	}
	if errExpected != errGotten {
		t.Errorf("Error message does not match. Got %q, want %q", errGotten, errExpected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Got an error but no error was expected")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Unexpected result - got %s want %s", got, want)
	}
}
