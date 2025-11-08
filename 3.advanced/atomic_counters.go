package advanced

import (
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func atomic_counterss() {
	var wg sync.WaitGroup
	numGoRoutines := 10

	counter := &AtomicCounter{}

	for range numGoRoutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}
	wg.Wait()
	println(counter.getValue())
}

/*
# What does "Atomic" mean?
	- indivisible
	- uniteruptable

# Why use atomic operations?
	- lost updates
	- inconsistent reads

# How atomic operations work?
	- lock-free
	- memory visibility

# Issues without atomic operations?
	- data race
	- inconsistent results

*/

/*
# Why use Atomic Counters?
	- performance
	- simplicity
	- concurrency

# Atomic operations in go
	- window duration
	- request counting
	- reset mechanism

# sync/atomic package:
	- atomic.AddInt32				/	atomic.AddInt64
	- atomic.LoadInt32				/	atomic.LoadInt64
	- atomic.StoreInt32				/	atomic.StoreInt64
	- atomic.CompareAndSwapInt32	/	atomic.CompareAndSwapInt64
*/
