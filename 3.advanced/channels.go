package advanced

import (
	"fmt"
)

func channels() {

	// variable := make (chan type)
	greeting := make(chan string)
	greetString := "Hello channel"

	// greeting <- greetString // blocking because it is continously trying to recieve values, it is ready to recieve continous flow of data

	go func() {
		greeting <- greetString // can also pass directly without declaring
		greeting <- "World"     // can also pass directly without declaring

		for _, e := range "abcde" {
			greeting <- string(e)
		}
	}()

	var receiver string

	// go func() {
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	// recieving values from channel
	receiver = <-greeting
	fmt.Println(receiver)
	receiver = <-greeting // part of main thread
	fmt.Println(receiver)

	for range 5 {
		temp := <-greeting
		fmt.Println(temp)
	}

	// time.Sleep(1 * time.Second) //use to avoid leaks when receiver is also in goroutine
	println("=== End ===")
}

/*
Channels are a way for goroutines to communicate and synchronize with each other

# Why use channels?
	-enable safe and efficient communication between concurent goroutines
	-help synchronize and manage the flow of data in concurrent programs
	* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

===	recievers will block all code and wait for all go routines to sleep before executing ahead ===

goroutines are non blocking, extracted away from main thread
channels are blocking

*/
