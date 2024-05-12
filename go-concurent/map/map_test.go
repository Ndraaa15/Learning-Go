package main

import (
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var data sync.Map
	var wg sync.WaitGroup
	var addToMap = func(wg *sync.WaitGroup, value int) {
		defer wg.Done()

		wg.Add(1)
		data.Store(value, value)
	}

	for i := 0; i < 100; i++ {
		go addToMap(&wg, i)
	}

	wg.Wait()

	data.Range(func(key, value interface{}) bool {
		t.Logf("key: %v, value: %v", key, value)
		return true
	})

}
