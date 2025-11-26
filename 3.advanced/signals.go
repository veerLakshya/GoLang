package advanced

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signals() {
	pid := os.Getpid()
	fmt.Println("Process pid: ", pid)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGHUP)

	go func() {
		sig := <-sigs
		fmt.Println("Recieved Signal: ", sig)
		done <- true
	}()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal.")
				return
			default:
				fmt.Println("Working...")
			}
			time.Sleep(time.Second)
		}

		// for sig := range sigs {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("Recieved SIGINT (Interrupt)")
		// 	case syscall.SIGTERM:
		// 		fmt.Println("Recieved SIGTERM (Terminate)")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("Recieved SIGHUP (Hangup)")
		// 	case syscall.SIGUSR1:
		// 		fmt.Println("Recieved SIGUSR1 (User Defined Signal 1)")
		// 		fmt.Println("User defined function is executed.")
		// 		continue
		// 	}
		// 	fmt.Println("Graceful exit.")
		// 	os.Exit(0)
		// }
	}()

	// simulate some work here
	// fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}

}

/*
# Why use signals?
	- graceful shutdown
	- resource cleanup
	- inter-process communication

# Signals in Unix like OS
	- SIGINT (Interrupt Signal)
	- SIGTERM (Terminate Signal)
	- SIGHUP (Hang Up Signal)
	- SIGKILL (Kill)

# Using the kill command
	- find the process id (pid)
	- send the signal
# Signal Types and Usage
	- Interrupts - SIGINT
	- Terminations - SIGTERM
	- Stop/Continue - SIGCONT, SIGSTOP


*/
