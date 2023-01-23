package golang_goroutine_udemy

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i < 1000; i++ {
		// these 1000 goroutine will cause race condition
		// because x is shared var and interacted with 1000 goroutines at once
		go func() {
			for j := 1; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
