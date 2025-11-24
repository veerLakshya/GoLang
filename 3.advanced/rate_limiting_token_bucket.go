package advanced

import (
	"log"
	"time"
)

type rateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

// WHY Empty Structs...?
// empty structs in go take up 0 bytes of memory
// signalling intent

func NewRateLimiter(rateLimit int, refillTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}
	for range rateLimit {
		rl.tokens <- struct{}{}
	}
	go rl.StartRefil()
	return rl
}

func (rl *rateLimiter) StartRefil() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			{
				select {
				case rl.tokens <- struct{}{}:
				default:
				}
			}
		}
	}
}

func (rl *rateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func tokenBucket() {
	rl := NewRateLimiter(6, time.Millisecond*10)

	for range 10 {
		if rl.allow() {
			log.Println("Request Allowed")
		} else {
			log.Println("Rate limit exceeded")
		}
		time.Sleep(200 * time.Millisecond)
	}
}

//

// 1 0 ms			first request allowed 			tokens left: 5
// 2 200 ms			second request allowed			tokens left: 4
// 3 400 ms			third request allowed			tokens left: 3
// 4 600 ms			fourth request allowed			tokens left: 2
// 5 800 ms			fifth request allowed			tokens left: 1
// 6 1000 ms		sixth request allowed			tokens left: 0		start refil fun executes and adds one more token, 1 token added
// 7 1200 ms		seventh request denied
// 8 1400 ms		eighth request denied
// 9 1600 ms		ninth request denied
// 10 1800 ms		tenth request denied
