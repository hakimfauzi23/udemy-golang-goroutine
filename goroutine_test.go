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
