package advanced

import (
	"fmt"
	"sync"
	"time"
)

type Workerr struct {
	ID   int
	Task string
}

// PerformTask simulates a worker performing a task
func (w *Workerr) PerformTaskk(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func wait_groups() {
	var wg sync.WaitGroup

	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Workerr{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTaskk(&wg)
	}

	// wait for all workers to finish
	wg.Wait()

	// Construction is finished
	fmt.Println("Construction finished")
}

// ===== EXAMPLE WITH CHANNELS =====
/*
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d starting. \n", id)
	time.Sleep(time.Second * 2)
	for task := range tasks {
		results <- task * 2
		fmt.Printf("WorkerID %d completed work %d. \n", id, task)
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 5
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i+1, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i + 1
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result: ", result)
	}
}
*/

// ===== BASIC EXAMPLE WITHOUT CHANNELS =====
/*
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting \n", id)

	time.Sleep(time.Second) // simulate some time for processing

	fmt.Printf("Worker %d finished. \n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i, &wg)
	}

	wg.Wait() // blocking mechanism
	fmt.Println("ALL Tasks completed")
}
*/

/*
# Why use Wait Groups?
	- synchronization
	- coordination
	- resource management

# Basic Operations -
	- Add(delta int)
	- Done()
	- Wait()

# Best Practices -
	- avoid Blocking Calls inside Goroutines
	- use `defer` to call `Done`
	- ensure proper use of `Add`
	- handle large number of goroutines
	- use `defer` for unblocking

# Common pitfalls -
	- mismatch between `Add` and `Done`
	- avoid creating deadlocks
*/
