package advanced

import (
	"log"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int           // max number of requests allowed in a window
	window    time.Duration // duration of a window
	resetTime time.Time     // time when the next window starts
}

func NewFixedWindowRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		limit:  limit,
		window: window,
	}
	return rl
}
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	curTime := time.Now()

	if curTime.After(rl.resetTime) {
		rl.resetTime = curTime.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func fixedwindowcounter() {
	rl := NewFixedWindowRateLimiter(5, time.Second*2)

	for range 10 {
		if rl.Allow() {
			log.Println("Request allowed")
		} else {
			log.Println("Request denied")
		}

		time.Sleep(time.Millisecond * 200)
	}
}

/*
# How it Works?
	- each window has a counter that tracks the number of requests
	- if the num of reqs in the current window is below the limit, the req is allowed and the counter is incremented
	- if the num of reqs reached the limit, subsequent requests in the same windo are denied
	- at the start of a new window, the counter is reset

# Key points
	- window duration
	- request counting
	- reset mechanism
*/
