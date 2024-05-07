package main

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	// Sync one memastikan bahwa function itu hanya akan dilakukan sekali
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			once.Do(OnlyOnce)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter : ", counter)
}
