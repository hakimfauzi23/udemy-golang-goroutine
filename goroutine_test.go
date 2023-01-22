package golang_goroutine_udemy

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // this will running asynchronous
	fmt.Println("Test@Run")

	time.Sleep(1 * time.Second) // will freeze 1 second
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i) // will print random because use goroutine
	}

	// time.Sleep(10 * time.Second)
}
