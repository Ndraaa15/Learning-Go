package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

// function di bawah berjalan secara synchronous
func TestCreateGoroutineSync(t *testing.T) {
	RunHelloWorld()

	fmt.Println("Hi World")
}

// function di bawah berjalan secara asynchronous
// dengan penambahan keyword go maka function RunHelloWorld() akan berjalan secara asynchronous yang berati function ini tidak akan ditunggu oleh proses setelahnya untuk di eksekusi
// namun hal ini terkedang menyebabkan function RunHelloWorld() tidak akan menampilkan outputnya karena program lebih dahulu selesai tanpa goroutine bisa mengakases function RunHelloWorld()
func TestCreateGoroutineAsync(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("Hi World")

	time.Sleep(2 * time.Second)
}

//-------------------------------------------------------------

func DisplayNumber(number int) {
	fmt.Println("Number : ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
