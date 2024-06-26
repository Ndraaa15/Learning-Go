package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var wg sync.WaitGroup

func WaitCondition(value int) {
	defer wg.Done()
	wg.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("value: ", value)
	cond.L.Unlock()

}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	wg.Wait()
}
