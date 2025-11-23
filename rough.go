package main

import (
	"fmt"
)

func fab(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	nums := 10
	c := make(chan int, nums)
	go fab(nums, c)
	for i := range c {
		fmt.Println(i)
	}
}
