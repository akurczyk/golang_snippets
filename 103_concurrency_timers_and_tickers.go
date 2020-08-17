package main

import (
	"fmt"
	"time"
)

func main() {
	//
	// Timer fires once
	//

	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	timer2.Stop() // Abort

	<-timer1.C // Wait for it



	//
	// Ticker fires periodically
	//
	ticker1 := time.NewTicker(1 * time.Second)
	ticker2 := time.NewTicker(2 * time.Second)

	go func() {
		for range ticker1.C {
			fmt.Println(time.Now())
		}
	}()

	time.Sleep(10 * time.Second)

	ticker1.Stop()
	ticker2.Stop()
}