package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go func() {
		time.Sleep(2 * time.Second)
		// Apabila channel yang sudah berisi data tidak ada yang mengambil akan terjadi error
		ch <- "Hello World"
		fmt.Println("Data succesfuly uploaded")
	}()

	result := <-ch
	fmt.Println(result)
}

func TestChannelAsParameter(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go getResponseChannel(ch)

	resp := <-ch

	fmt.Println(resp)
}

func getResponseChannel(ch chan string) {
	ch <- "Indra"
}

func TestChannelInOut(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go inChannel(ch)
	go outChannel(ch)
	// Spawn two goroutine

	// go func() {
	// 	inChannel(ch)
	// 	outChannel(ch)
	// }()
	// code above just only spawning one goroutine because it using anonymous function that include function inChannel and outChhanel so the process it will be inChananel first but got error because no one take the value in channel
	time.Sleep(3 * time.Second)
}

func inChannel(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Indra"
}

func outChannel(ch <-chan string) {
	str := <-ch
	fmt.Println(str)
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan string, 3)

	ch <- "Gede"
	ch <- "Indra"
	ch <- "Adi"

	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch))

	defer close(ch)
}

func TestRangeChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- "Perulangan ke " + strconv.Itoa(i)
		}
	}()

	for data := range ch {
		fmt.Println(data)
	}
}

func TestChannnnn(t *testing.T) {
	ch := make(chan string)

	go func() {
		str := returnFunc()
		ch <- str
	}()

	a := <-ch
	fmt.Println(a)
}

func returnFunc() string {
	return "indra"
}
