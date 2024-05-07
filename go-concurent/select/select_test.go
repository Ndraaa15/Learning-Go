package main

import (
	"fmt"
	"testing"
	"time"
)

func getResponseChannel(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Indra"
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go getResponseChannel(ch1)
	go getResponseChannel(ch2)

	// dalam hal ini select digunakan untuk mendapatkan data dari channel yang pertama kali memberikan data
	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data dari ch 1 : ", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data dari ch 2 : ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go getResponseChannel(ch1)
	go getResponseChannel(ch2)

	// dalam hal ini default select akan mengeksekusi kode di bloknya selama belum ada data di dalam channel tersebut (di dalam select)
	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data dari ch 1 : ", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data dari ch 2 : ", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
