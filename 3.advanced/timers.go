package advanced

import (
	"fmt"
	"time"
)

func timers() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("timer1 expired.")
	case <-timer2.C:
		fmt.Println("timer2 expired.")
	}

}

// === SCHEDULING DELAYED OPERATIONS ===
/*
func main() {
	timer := time.NewTimer(2 * time.Second) // non blocking timer starts

	go func() {
		<-timer.C
		fmt.Println("Delayed operation executed")
	}()

	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second) // blocking timer starts
}
*/

// === TIMEOUT ===
/*
func longRunningOperation() {
	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	timeout := time.After(2 * time.Second)
	done := make(chan bool)

	go func() {
		longRunningOperation()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("Operation timed out")
	case <-done:
		fmt.Println("Operation Completed")
	}
}
*/

// === BASIC Timer USE ===
/*
func main() {
	timer := time.NewTimer(2 * time.Second)

	stopped := timer.Stop()

	if stopped {
		fmt.Println("timer stopped")
	}

	timer.Reset(time.Second)
	fmt.Println("Timer resetted")
	<-timer.C // blocking in nature
	fmt.Println("Timer expired")
}
*/

/*
# Why use timers?
	- timeouts
	- delays
	- periodic tasks

# The `time.Timer` type
	- Creating a Timer
	- Timer channel
	- Stopping a Timer

# Use cases -
	- implementing timeouts
	- scheduling delayed operations
	- periodic tasks
	- handle large numbers of goroutines
	- use `defer` for unlocking

	* avoid resource leaks
	* combine with channels
	* use `timer.After` for simple timeouts

*/
