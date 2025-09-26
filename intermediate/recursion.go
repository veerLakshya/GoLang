package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
func recurse() {
	fact := factorial(5)
	sum := sumOfDigits(fact)
	fmt.Println(sum)
}
