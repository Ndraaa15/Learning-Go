package testing

import (
	"fmt"
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

	time.Sleep(5 * time.Second)
}
