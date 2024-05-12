package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("x:", x)
}
