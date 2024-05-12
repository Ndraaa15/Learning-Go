package timer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Current time: ", time.Now())

	time := <-timer.C
	fmt.Println("Timer expired at: ", time)
}

func TestTimerr(t *testing.T) {
	// No need duration
	channel := time.After(5 * time.Second)
	fmt.Println("Current time: ", time.Now())

	time := <-channel
	fmt.Println("Timer expired at: ", time)
}

func TestTimerrr(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Timer expired at: ", time.Now())
		wg.Done()
	})
	wg.Wait()
	fmt.Println("Current time: ", time.Now())
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println("Ticker at: ", time)
	}
}

func TestTickerr(t *testing.T) {
	ticker := time.Tick(2 * time.Second)

	for time := range ticker {
		fmt.Println("Ticker at: ", time)
	}
}
