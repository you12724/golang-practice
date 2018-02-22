package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	go func() {
		Withdraw(100)
		done <- struct{}{}
	}()

	go func() {
		_, ok := Withdraw(1000)
		if ok {
			t.Error("can't withdraw")
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done
	<-done

	if got, want := Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
