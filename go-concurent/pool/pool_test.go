package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	// New akan digunakan apabila ketika pool kosong dan terdapat goroutine yang inget mengakses pool, maka akan diberikan data new tersebut
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	pool.Put("Indra")
	pool.Put("Adi")
	pool.Put("Brata")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
			wg.Done()
		}()
	}

	wg.Wait()
}
