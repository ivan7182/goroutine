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
