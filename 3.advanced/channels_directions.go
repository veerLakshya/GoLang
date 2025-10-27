package advanced

import "fmt"

func channelDirections() {

	ch := make(chan int)

	// SEND only Channel
	producer(ch)

	// RECIEVE only channel
	consumer(ch)

}

// SEND only
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// RECIEVE only
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Recieved: ", value)
	}
}

/*
# Why are channel directions important?
	- improve code clarity and maintainability
	- prevent unintended operations on channels
	- enhance type safety by clearly defining the channel's purpose

# Basic concepts of channel directions-
	- unidirectional channels
	- send-only channels
	- recieve-only channels
	- testing and debugging
# Defining channel directions in function signatures-
	- send-only parameters(func produceData(ch chan<-int))
	- receive-only parameters(func consumeData(ch <-chan int))
	- bidirectionaly channels(func bidirectional(ch chan int))
*/
