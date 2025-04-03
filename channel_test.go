package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "ivan satria"
		fmt.Println("selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// channel parameter

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "ivan satria"
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// In Out Channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "ivan satria"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {

		channel <- "Ivan"
		channel <- "satria"
		channel <- "joko"

	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()

	time.Sleep(5 * time.Second)
	fmt.Println("selesai")

}

//Channel Range

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("selesai")

}

//Select Channel

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

//Default Channel

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
