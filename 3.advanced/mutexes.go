package advanced

import (
	"fmt"
	"sync"
)

func mutexess() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()

	fmt.Println("value", counter)
}

/*
type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	counter := &counter{}

	numGoroutines := 10

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// counter.count++ -> WRONG since there is no check and multiple goroutines might update it at the same time
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.getValue())

}
*/

/*
# Why Mutual Exclusion is Important?
	- Data integrity
	- Consistency
	- Safety

# How mutual exclusion is achieved?
	- locks (Mutexes)
	- semaphores
	- monitors
	- critical sections

# Why the are often used in structs?
	- encapsulation
	- convenience
	- readibility
*/

/*
# Why use mutexes?
	- Data integrity
	- synchronization
	- avoid race conditions

# Basic Operations
	- Lock()
	- Unlock()
	- TryLock()

# Mutex and Performance
	- contention
	- granularity

# Best Practices -
	- minimize lock duration
	- avoid nested locks
	- prefer sync.RWMutex for Read-Heavy Workloads
	- check for deadlocks
	- use `defer` for unblocking

# Common pitfalls -
	- deadlocks
	- performance issues
	- starvation

*/
