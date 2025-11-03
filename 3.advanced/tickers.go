package advanced

import (
	"fmt"
	"time"
)

// ===== TICKER WITH TIMER PATTERN =====
func tickerss() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)

	defer ticker.Stop()

	for {
		select {

		case tick := <-ticker.C:
			fmt.Println("Tick at: ", tick)
		case <-stop:
			fmt.Println("Stopping Ticker.")
			return
		}
	}
}

// ===== BASIC PERIODIC TASK EXECUTION, scheduling logging, periodic tasks, polling tasks =====
/*
func periodicTask() {
	fmt.Println("Performing periodic task at: ", time.Now())
}

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			periodicTask()
		}
	}
}
*/

// func main() {
// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at: ", tick)
// 	// }
// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// }

/*
# Why use tickers?
	- consistency
	- simplicity
	- handling periodic tasks
# Best Pracitcies
	- stop tickers when no longer needed
	- avoid blocking operations
	- handle ticker stopping gracefully
*/
