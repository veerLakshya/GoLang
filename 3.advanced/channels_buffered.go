package advanced

import "time"

func bufferedChannels() {
	// === BLOCKING ON RECEIVE WHEN THE BUFFER IS EMPTY ===

	ch := make(chan int, 2)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- 1
	// }()
	// println("value: ", <-ch) // waits for all goroutines to end if buffer is empty

	ch <- 1
	ch <- 2
	println("receiving from buffer")
	go func() {
		time.Sleep(2 * time.Second)
		println("Recieved: ", <-ch) // executes right to left (ends <<<<----- starts) for all programming language
	}()

	println("Blocking starts")
	ch <- 3
	println("Blocking ends")

}

// func main() {
// ==== BLOCKING ONLY WHEN BUFFER IS FULL ===
// 	ch := make(chan int, 2)

// 	ch <- 1
// 	ch <- 1
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		println("recieved: ", <-ch)
// 	}()
// 	ch <- 3 // blocks because  buffer full and thats why print statement in goroutine executes

// 	println("Buffered channels")
// }

/*
#Why use unbuffered channels?
	-asynchronous communication
	-load balancing
	-flow control
	-only blocks when capacity is full
	-if capacity of buffer is full, it waits for all goroutines to end before executing that thread(since a reciever may recieve from it there and making space more space there)

#Creating buffered channels-
	-make (chan type, capacity)
	-buffer capacity

#Key things of buffered channels-
	-blocking behaviour
	-non-blocking operations
	-impact on performance

#Best Practices while using buffered channels-
	-avoid over-buffering
	-graceful shutdown
	-monitoring buffer ussage

sync	 	- use unbuffered channels	- 	blocks until receiver receives 			- used in direct handoff
async		- use buffered channels 	- 	blocks when capacity is full or empty	- used in async messaaging and queing
*/
