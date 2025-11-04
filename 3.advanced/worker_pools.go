package advanced

import (
	"fmt"
	"time"
)

type tickerRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate processing of tickets requests
func ticketProcessor(requests <-chan tickerRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d tickets of personID %d with total cost %d \n", req.numTickets, req.personID, req.cost)

		// simulate processing time of tickets
		time.Sleep(time.Second * 2)
		results <- req.personID
	}
}

func workerpoolss() {
	numRequests := 10
	price := 5
	ticketRequests := make(chan tickerRequest, numRequests)
	ticketResults := make(chan int)

	// start 3 workers
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {
		ticketRequests <- tickerRequest{personID: (i + 1), numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully! \n", <-ticketResults)
	}
}

/* ====== BASIC WORKER POOL PATTERN
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d \n", id, task)

		// simulate some work
		time.Sleep(time.Second)
		results <- task * 2
	}
}

func main() {
	numWorkers := 3
	numJobs := 10

	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create workers
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	// Send tasks to the tasks channel
	for i := range numJobs {
		tasks <- i
	}

	close(tasks)

	for range numJobs {
		result := <-results
		fmt.Println("Result: ", result)
	}
}
*/

/*

# Why use worker pools?
	- resource management
	- task distribution
	- scalibility

# Conceptual Model -
	- tasks
	- workers
	- queues

# Implementation Steps-
	- create a task channel
	- create worker goroutines
	- distribute tasks
	- graceful shutdown

# Advanced Worker Pool Patterns -
	- dynamic worker pools
	- task prioritization
	- error handling
	- worker pool with task prioritizing

# Best Practices -
	- limit the number of workers
	- handle worker lifecycle
	- implement timeouts
	- monitor and log
*/
