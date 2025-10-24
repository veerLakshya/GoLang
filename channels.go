package main

import "fmt"

func main() {

	// variable := make (chan type)

	greeting := make(chan string)
	greetString := "Hello channel"

	// greeting <- greetString // blocking because it is continously trying to recieve values, it is ready to recieve continous flow of data

	go func() {
		greeting <- greetString // can also pass directly without declaring
	}()

	// recieving value from channel
	receiver := <-greeting
	fmt.Println(receiver)
}

/*
Channels are a way for goroutines to communicate and synchronize with each other

# Why use channels?
	-enable safe and efficient communication between concurent goroutines
	-help synchronize and manage the flow of data in concurrent programs
	* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

goroutines are non blocking, extracted away from main thread
channels are blocking

*/
