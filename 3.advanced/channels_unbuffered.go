package advanced

import "time"

func unbufferedChannel() {
	ch := make(chan int)

	var re int
	go func() {
		time.Sleep(2 * time.Second)
		re = <-ch

	}()
	ch <- 1
	// re = <-ch // recievers will block all code and wait for all go routines to sleep before executing ahead
	println("asdf")
	println(re)
}

/*
-Channel is kind of a storage
	-By default a unbuffered channel is created using our `name := make (chan type)` syntax
	*Unbuffered channels cannot hold values thats why they require immediate receiver
	-Buffered channel - a channel with an associated storage(limitted space)
	-Buffered channels wait for all go routines to finish

-unbuffered channels allow async communication
-buffered channels dont block until buffer is full
-buffered channels also help in load balancing, handling burst of data without immediate synchronization and hence allowing better flow control
-unbuffered channels require immediate receivers
*/
