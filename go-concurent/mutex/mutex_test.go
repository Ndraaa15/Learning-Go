package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x = 0
	var mutex sync.Mutex
	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				// Dengan menggunakan locking mutex maka hanya ada 1 goroutine yang bisa mengakses shared var
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance float64
}

func (acc *BankAccount) AddBalance(amount float64) {
	acc.RWMutex.Lock()
	acc.Balance += amount
	acc.RWMutex.Unlock()
}

func (acc *BankAccount) GetBalance() float64 {
	acc.RWMutex.RLock()
	balance := acc.Balance
	acc.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	accBank := BankAccount{}
	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				accBank.AddBalance(1)
				fmt.Printf("Salo : %.0f\n", accBank.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
}
