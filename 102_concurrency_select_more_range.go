package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	chan3 := make(chan string)



	//
	// Get first available message locking until one arrives
	//

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- ":-)"
	}()

	select {
	case msg1 := <-chan1:
		fmt.Println("Got message from channel no 1:", msg1)
	case msg2 := <-chan2:
		fmt.Println("Got message from channel no 2:", msg2)
	case msg3 := <-chan3:
		fmt.Println("Got message from channel no 3:", msg3)
	}



	//
	// Get first available or execute the code from "default" section
	//

	select {
	case msg1 := <-chan1:
		fmt.Println("Got message from channel no 1:", msg1)
	case msg2 := <-chan2:
		fmt.Println("Got message from channel no 2:", msg2)
	case msg3 := <-chan3:
		fmt.Println("Got message from channel no 3:", msg3)
	default:
		fmt.Println("No messages :-(")
	}



	//
	// Get first available or give up after a while
	//

	select {
	case msg1 := <-chan1:
		fmt.Println("Got message from channel no 1:", msg1)
	case msg2 := <-chan2:
		fmt.Println("Got message from channel no 2:", msg2)
	case msg3 := <-chan3:
		fmt.Println("Got message from channel no 3:", msg3)
	case <-time.After(1 * time.Second):
		fmt.Println("No messages :-(")
	}



	//
	// Close and more
	//

	type params struct {
		a, b int
	}

	inChannel := make(chan params)
	outChannel := make(chan int)

	go func(in <-chan params, out chan<- int) {
		var params params
		var more bool = true

		for {
			params, more = <-in
			if !more { break }
			out <- params.a * params.b
		}

		fmt.Println("Closing side thread...")
	}(inChannel, outChannel)

	for i := 0; i < 10; i++ {
		inChannel <- params{i, 10}
		fmt.Println(<-outChannel)
	}
	close(inChannel)

	time.Sleep(1 * time.Second)
	fmt.Println()



	//
	// Close and range
	// It does exactly the same as example above.
	//

	inChannel = make(chan params)
	outChannel = make(chan int)

	go func(in <-chan params, out chan<- int) {
		for params := range in {
			out <- params.a * params.b
		}

		fmt.Println("Closing side thread...")
	}(inChannel, outChannel)

	for i := 0; i < 10; i++ {
		inChannel <- params{i, 10}
		fmt.Println(<-outChannel)
	}
	close(inChannel)

	time.Sleep(1 * time.Second)
	fmt.Println()
}
