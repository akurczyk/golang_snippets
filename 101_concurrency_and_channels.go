package main

import (
	"fmt"
	"time"
)

func sideThread() {
	time.Sleep(1 * time.Second)
	fmt.Println("Side thread")
}

func main() {
	//
	// Basic example
	//

	go sideThread()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Another side thread")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main thread")

	time.Sleep(1 * time.Second)
	fmt.Println()



	//
	// Channels
	//

	channel := make(chan string) // Channel with no buffer, it locks if there is nothing that tries to get the message
	// channel := make(chan string, 1) // It can hold up to one message
	// channel := make(chan string, 10) // It can store 10 messages in its buffer

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "ping"
		time.Sleep(1 * time.Second)
		channel <- "ping2"
	}()

	message := <-channel // It locks until the message is available and stores it in the variable
	fmt.Println(message)

	<-channel // It locks and discards the message - useful for synchronization, mostly used with bool

	time.Sleep(1 * time.Second)
	fmt.Println()



	//
	// One way channels passing
	//
	type params struct {
		a, b int
	}

	inChannel := make(chan params)
	outChannel := make(chan int)

	go func(in <-chan params, out chan<- int) {
		params := <-in
		result := params.a * params.b
		out <- result
	}(inChannel, outChannel)

	inChannel <- params{10, 20}
	fmt.Println(<-outChannel)
	fmt.Println()
}
