package advanced

import (
	"fmt"
)

func multiplexing() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // signals that no more data will be sent
	}()

	for range ch {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel Closed")
				// clean up activities
				return
			}
			fmt.Println("Recieved: ", msg)
		}
	}
}

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	go func() {
// 		time.Sleep(time.Second)
// 		ch1 <- 1
// 	}()
// 	go func() {
// 		time.Sleep(time.Second)
// 		ch2 <- 22
// 	}()
// 	// if default is there, it is executed asap and doesnt wait for case 1 or 2
// 	for range 2 {
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("Recieved from ch1: ", msg)
// 		case msg := <-ch2:
// 			fmt.Println("Recieved from ch2: ", msg)
// 			// default:
// 			// 	fmt.Println("No channels ready")
// 		}
// 	}
// }

// SELECT WITH TIMEOUT
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Recieved: ", msg)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("Timeout.")
// 	}
// }

/*
# What is multiplexing?
	- combining multiple input streams into a single "listener" that can respond to whichever event happens first

# Why use mulitplexing?
	- manage mulitple concurrent operations within same go routine
	- non-blocking operations
	- timeouts and cancellations

# Best practices for using `select`-
	- avoid busy waiting
	- handling deadlocks
	- readability and maintainability
	- testing and debugging


# MULTIPLEXING LOOP PATTERN- to continously listen to some channels
	for {
		select {
		case msg := <-ch1:
			fmt.Println("From ch1:", msg)
		case msg := <-ch2:
			fmt.Println("From ch2:", msg)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout — no messages received")
			return
		}
	}

# FAN-IN Pattern - merge multiple channels into one output channel
	func fanIn(ch1, ch2 <-chan string) <-chan string {
		out := make(chan string)
		go func() {
		for {
			select {
				case msg := <-ch1:
					out <- msg
				case msg := <-ch2:
					out <- msg
				}
			}
		}()
		return out
	}

	merged := fanIn(ch1, ch2)
	for msg := range merged {
		fmt.Println(msg)
	}


#
Multiplexing	-	Listening to multiple channels at once				-	select
Demultiplexing	-	Sending data from one source to multiple channels	-	Multiple sends (ch1 <- data, ch2 <- data)
Fan-in			-	Combining multiple channels into one				-	Merging with select
Fan-out			-	Multiple goroutines reading from one channel		-	Parallel processing



*/
