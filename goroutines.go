package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() { // this fn is running in main thread
	fmt.Println("NUMBER OF CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(2))

	go sayHello() // go keyword takes this fn to a new thread

	go printNumbers()
	go printLetters()

	//handling errors or return values with goroutines
	var err error
	go func() {
		err = doWork()
	}()
	time.Sleep(time.Second * 2)
	if err != nil {
		log.Fatalf("error doing work: %v", err)
	}
}

func sayHello() {
	time.Sleep(time.Second * 1)
	fmt.Println("Hello from goroutine")
}

func printNumbers() {
	for i := 0; i < 25; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}
func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occurend doing work")
}

/*
# What are goroutines?
	-goroutines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finised/ready to return any value
	-they do not stop the program flow and are non blacking

# Why use goroutines?
	1- simplify concurrent programming
	2- efficiently handle concurrent tasks
	3- a way to perform tasks concurrently without manually managing threads

# Goroutine scheduling in go -
	-managed by the go runtime scheduler
	-uses m:n scheduling model (m goroutines on n threads)
	-efficient multiplexing
	-cant expect any order of fulfilling of created routines

# Pitfalls -
	-avoiding goroutine leaks
	-limiting goroutine creation
	-proper error handling
	-synchronization
*/
