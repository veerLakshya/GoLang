package advanced

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWORK(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled: ", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func contexttt() {
	rootCtx := context.Background() // root context from which other contexts are derived

	ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second) // context timer starts here, two seconds to use this
	defer cancel()

	// Can also use .WithCancel and cancel after a specific time using goroutine

	ctx = context.WithValue(ctx, "requestID", "asdf23f1fe1313")

	go doWORK(ctx)
	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID: ", requestID)
	} else {
		fmt.Println("No request ID found.")
	}

	logWithContext(ctx, "loggin with ctx")
}

func logWithContext(ctx context.Context, message string) {
	requestIDVal := ctx.Value("requestID")
	log.Printf("RequestID: %v - %v", requestIDVal, message)
}

/*
func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation canceled."
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func main() {
	ctx := context.TODO()

	result := checkEvenOdd(ctx, 3)
	fmt.Println("Result with ctx.TODO: ", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	result = checkEvenOdd(ctx, 6)
	fmt.Println("Result from timeout context: ", result)

	time.Sleep(2 * time.Second)
	result = checkEvenOdd(ctx, 12)
	fmt.Println("Result after timeout: ", result)
}
*/

// === Difference between .TODO and .BACKGROUND ===
/*
func main() {
	todoContext := context.TODO()
	contextBkg := context.Background()

	ctx := context.WithValue(todoContext, "name", "Lakshya")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	ctx2 := context.WithValue(contextBkg, "city", "Bareilly")
	fmt.Println(ctx2)
	fmt.Println(ctx2.Value("city"))
}
*/

/*

# Why use context?
	- cancellation
	- timeouts
	- values

# Basic concepts
	- context creation
		- context.Backgroud()
		- context.TODO()
	- context Hierarch(How contexts are created and derived)
		- context.WithCancel()
		- context.WithDeadline()
		- context.WithTimeout()
		- context.WithValue()

# Best Practices
	- avoid storing contexts in structs
	- propogating ctx correctly
	- handling ctx values
	- handling ctx cancellation
	- avoid creating contexts in loops
	- not misusing context values

*/
