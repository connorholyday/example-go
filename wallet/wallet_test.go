package wallet

import (
    "testing"
    "fmt"
)

func TestWallet(t *testing.T) {

    t.Run("deposits money", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        got := wallet.Balance()
        want := Bitcoin(10)

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    })

    t.Run("withdraws money", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        wallet.Withdraw(Bitcoin(3))

        got := wallet.Balance()
        want := Bitcoin(7)

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    })

    t.Run("errors on empty withdrawal", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        err := wallet.Withdraw(Bitcoin(20))

        if err == nil {
            t.Errorf("wanted an error but didn't get one")
        }
    })
}

func ExampleWallet() {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(10))

    fmt.Println(wallet.Balance())
    // Output: 10 BTC
}
