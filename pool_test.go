package golang_goroutine_udemy

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New" // will set default data, if pool is empty.
		},
	}

	pool.Put("localhost")
	pool.Put("mysqli")
	pool.Put("192.168.1.1")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("End Of Test Pool")

}
