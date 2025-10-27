package advanced

import (
	"fmt"
	"time"
)

// func main() {

// 	done := make(chan struct{})

// 	go func() {
// 		println("Working..")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	<-done
// 	println("Finished.")

// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9
// 		time.Sleep(time.Second)
// 		println("value sent")
// 	}()

// 	value := <-ch
// 	println(value)
// }

// ==== SYNCHRONIZING MULTIPLE GOROUTINES and ensuring that all goroutines are completed ===
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines {
// 		time.Sleep(2 * time.Second)

// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working\n", id)
// 			time.Sleep(time.Second)
// 			done <- id // Sending singnal of Completion
// 		}(i)
// 	}

// 	for range 2 {
// 		<-done // wait for each goroutine to finish
// 	}
// 	println("== all go routines executed ==")
// }

// === SYNCHRONIZING DATA EXCHANGE ===
func synchronization() {
	data := make(chan string)
	go func() {
		for i := range 5 {
			data <- fmt.Sprintf("Hello %d", i)
			time.Sleep(time.Second)
		}
		close(data) // used to close channel once use is done so that reciever knows when to stop listening
	}()

	// Loops over only on active channel, creates reciever each time and stops once channel is closed, deadlocks at end when incoming values stop if channel is not closed
	for value := range data {
		fmt.Println("recieved value: ", value, " : ", time.Now())
	}
}

/*
#Why is channel synchronization important?
	- ensure that data is properly exchanged between goroutines
	- coordinates the execution flow to avoid race conditions and ensure predictable behaviour
	- helps manage the lifecycle of goroutines and the completion of tasks

#Common things to take care-
	- deadlocks
	- closing channels
	- avoiding unnecessary blocking
	- sequence of finishing can be different
*/
