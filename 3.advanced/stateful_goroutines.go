package advanced

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	count int
	ch    chan int
}

func (w *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current count:", w.count)
			}
		}
	}()
}

func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

func statefulWorkerExample() {
	stWorker := &StatefulWorker{
		ch: make(chan int),
	}

	stWorker.Start()

	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(time.Millisecond * 500)
	}
}

/*
maintains and updates its own state

# Why use stateful goroutines?
	- state management
	- concurrency
	- task execution

# Key Concepts of Stateful Goroutines
	- state preservation
	- concurrency management
	- lifecycle management

# Handling Concurrency and Synchronization
	- mutexes and locks
	- atomic operations
	- channels

# Common Use Cases -
	- task processing
	- stateful services
	- data stream processing

# Best Practices
	- Encapsulate state
	- synchronize access
	- monitor and debug
*/
