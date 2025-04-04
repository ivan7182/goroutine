package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1) // Menambah jumlah goroutine yang harus ditunggu
		go func() {
			defer wg.Done() // Memberitahu bahwa goroutine ini telah selesai
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	wg.Wait() // Menunggu semua goroutine selesai sebelum melanjutkan
	fmt.Println("counter =", x)

	// Menambahkan assertion untuk memastikan hasilnya sesuai dengan ekspektasi
	expected := 1000 * 100
	if x != expected {
		t.Errorf("Expected counter to be %d, but got %d", expected, x)
	}
}

type BankAccount struct {
	RwMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RwMutex.Lock()
	account.Balance = account.Balance + amount
	account.RwMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RwMutex.RLock()
	balancee := account.Balance
	account.RwMutex.RUnlock()
	return balancee
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total balance ", account.GetBalance())
}

//Deadlock

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amout int) {
	user1.Lock()
	fmt.Println(" lock user1", user1.Name)
	user1.Change(-amout)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println(" lock user2", user2.Name)
	user2.Change(amout)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Ivan",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "satria",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)

}
