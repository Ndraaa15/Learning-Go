package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var x = 0
	// nilai yang tidak konsisten disebabkan oleh goroutine yang mengakses shared var secara bersamaan sehingga tidak terjadi perubahan yang sesuai
	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}
