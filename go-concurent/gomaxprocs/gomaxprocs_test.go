package gomaxprocs

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	numCpu := runtime.NumCPU()
	fmt.Println("Num CPU: ", numCpu)

	numThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Num Thread: ", numThread)

	numGoroutine := runtime.NumGoroutine()
	fmt.Println("Num Goroutine: ", numGoroutine)
	// 2 Goroutine karena terdapat garbarge collector
}
