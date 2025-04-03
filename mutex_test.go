package main

import (
	"fmt"
	"sync"
	"testing"
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
