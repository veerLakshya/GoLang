package advanced

import "fmt"

func producer2(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}
func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

func closingg() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer2(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println("123: ", val)
	}
}

// === RANGE OVER CLOSED CHANNEL ===
/*
func main() {
	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("recieved val: ", val)
	}
}
*/

// === RECEIVING FROM  A CLOSED CHANNEL ===
/*
func main() {
	ch := make(chan int)
	close(ch)

	val, ok := <-ch
	if !ok {
		fmt.Println("Channel is closed.")
		return
	}
	fmt.Println(val)
}
*/

// === SIMPLE CLOSING EXAMPLE ===
/*
func main() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
*/

/*

# Why close channels?
	- signal completion
	- prevent resource leaks

# Best practices for closing channels-
	- close channels only from the sender
	- avoid closing channels more than once (goes in runtime panic)
	- avoid closing channels from multiple goroutines

# Common patters for closing channels -
	- pipeline pattern
	- worker pool pattern

# Debugging and troubleshootin channel closures -
	- identify closing channel errors
	- use `sync.WaitGroup` for coordination

*/
