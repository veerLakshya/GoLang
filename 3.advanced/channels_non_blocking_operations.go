package advanced

import (
	"fmt"
	"time"
)

func nonBlockingOperations() {

	/*
		// === NON BLOCKING RECEIVE OPERATION

		ch := make(chan int)

		select {
		case msg := <-ch:
			fmt.Println("Received: ", msg)
		default:
			fmt.Println("NO messages available.")
		}

		// === NON BLOCKING SEND OPERATION
		select {
		case ch <- 1:
			fmt.Println("Message sent.")
		default:
			fmt.Println("Channel is not ready to recieve.")
		}
	*/

	// === NON BLOCKING OPERATIONS in real time systems
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data recieved: ", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(1 * time.Second)
	}
	quit <- true

}

/*

# Why use non-blocking operations?
	- avoid deadlocks
	- improve efficiency
	- enhance concurrency

# Practices-
	- avoid busy waiting
	- handle channel closure properly
	- combine with context for cancellations
	- ensure channel capacity management

*/
