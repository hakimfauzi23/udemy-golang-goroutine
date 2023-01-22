package golang_goroutine_udemy

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	channel := make(chan string) // make a channel
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hanif Fauzi Hakim"
		fmt.Println("End sending data to channel")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// channel can be as parameter of function
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hanif Fauzi Hakim"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string) // make a channel
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

// this func use for sending data to channel only
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hanif Fauzi Hakim"
}

// this func use for retrieve data from channel only
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {

	// buffered is to avoid blocking when channel receiver is slow.
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Hanif"
		channel <- "Fauzi"
		channel <- "Hakim"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
