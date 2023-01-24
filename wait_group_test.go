package golang_goroutine_udemy

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // important!

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	// WaitGroup is to use when there are many goroutine and
	// don't know exact time to finish all goroutine
	group.Wait()
	fmt.Println("End of goroutine")
}
