package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance float64
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) ChangeBalance(amount float64) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount float64) {
	user1.Lock()
	fmt.Println("Lock User1 : ", user1.Name)
	user1.ChangeBalance(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User2 : ", user2.Name)
	user2.ChangeBalance(amount)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Indra",
		Balance: 10000,
	}

	user2 := UserBalance{
		Name:    "Bayu",
		Balance: 50000,
	}

	go Transfer(&user1, &user2, 5000) // groutine 1
	go Transfer(&user2, &user1, 5000) // goroutine 2

	// Proses deadlock ini terjadi karena goroutine 1 sedang me-locking user 1 (user 2 di goroutine 2)
	// dan goroutine 2 sedang melakukan locking pada user 1 (user 2 di goroutine 1)
	// sehingga ketika setiap goroutine ingin mengakses user tersebut terjadi deadlock (saling menunggu)

	time.Sleep(10 * time.Second)
	fmt.Println("user1.Name : ", user1.Name, "user1.Balance : ", user1.Balance)
	fmt.Println("user2.Name : ", user2.Name, "user2.Balance : ", user2.Balance)
}
