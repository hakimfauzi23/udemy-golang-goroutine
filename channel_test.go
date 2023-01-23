package golang_goroutine_udemy

import (
	"fmt"
	"strconv"
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

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 25; i++ {
			channel <- "Loops : " + strconv.Itoa(i)
		}
		close(channel)
	}()

	// with for range is useful if the channel if many.
	for data := range channel {
		fmt.Println("Retrieve data", data)
	}

	fmt.Println("End of channel")
}

func TestSelectChannel(t *testing.T) {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	/*
		With select channel it's possible to get data more than
		one channel advanced more than for range channel.
	*/
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Retrieving data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Retrieving data from channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	/*
		With select channel it's possible to get data more than
		one channel advanced more than for range channel.
	*/
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Retrieving data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Retrieving data from channel 2", data)
			counter++
		default:
			fmt.Println("Waiting for the data") // default select is to avoid deadlock.
		}

		if counter == 2 {
			break
		}
	}
}
