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

func TestChannel(t *testing.T) {

	// Buffered Channel
	// Channel that can store limited value
	ch1 := make(chan int, 2)

	// Unbuffered Channel
	ch2 := make(chan int)

	/*
		ch <- 1
		n := <-ch
		fmt.Println(n)

		Code above will cause deadlock because channel is not buffered
		And also when we send or recieve data it must be at the same time so wee need to use goroutine
		Channel is like portal so after we go into a portal and simustaneously we go out from the portal
	*/

	/*
		go func() {
			ch <- 1
		}()
		n := <-ch
		println(n)

		Code above will work properly because already using goroutine
		This make the channel can send and recieve data at the same time (simultaneously)
		This is called asynchronous
	*/

	ch1 <- 1
	num1 := <-ch1
	println(num1)

	go func() {
		ch2 <- 2
	}()

	num2 := <-ch2
	print(num2)
}
