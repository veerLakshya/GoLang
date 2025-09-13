package main

import (
	"fmt"
)

func main() {
	// ... create an array which is good enough to hold the data we are initializing
	grades := [...]int{1, 2, 3}
	fmt.Println(grades)

	var students [3]string
	students[0] = "Alice"
	students[1] = "Bob"
	fmt.Printf("students: %v\n", students)
	println(len(students)) // get length/size of array

	var iMatrix [3][3]int
	iMatrix[0] = [3]int{1, 2, 3}
	iMatrix[1] = [3]int{1, 2, 3}
	iMatrix[2] = [3]int{1, 2, 3}
	fmt.Println(iMatrix)

	a := [...]int{1, 2, 3}
	// b := a  // copying arrays creates new array
	b := &a // now b is pointing to the same data as in a
	b[1] = 5
	fmt.Println(a, b)

	//*Slice
	sl := []int{1, 2, 3}
	fmt.Println(sl, len(sl))
	fmt.Println(cap(a))
	// Naturally reference types
	/*
		p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		q := a[:]   // slice of all elements
		r := a[3:]  // slice from 4th element to end
		s := a[:6]  // slice of first 6 elements
		t := a[3:6] // slice the 4th, 5th, 6th element
	*/

	u := make([]int, 3, 100) // takes 2 or 3 arguments, type of object we want to make, length of slice, capacity of slice
	fmt.Println(u)

	p := []int{}
	fmt.Println(p)      // Output: []
	fmt.Println(len(p)) // Output: 0
	fmt.Println(cap(p)) // Output: 0

	p = append(p, 2, 3, 4, 5, 6)

	fmt.Println(p)

	// concatenation

	p = append(p, []int{7, 8, 9}...)

	// popping first, last
	q := a[:len(a)-1]
	println(q)

	/* Summary
	Arrays
		-collection of items with same types
		-fixed size
		-declaration styles-
			a := [3]int{1,2,3}
			a := [...] int{1,2,3}
			var a [3]int
		-access via zero-based indexes
			a[1], a[0]
		-len fn return size of array
		-copies refer to different underlying data

	Slices
		-Backed by arrays
		-creation stlyes-
		-slice existing array or slice
		-literal style
		-via make function
			- a := make([]int, 10)
			- a := make([]int, 10 ,100) size = 10, cap = 100
		- len fn gives length of slice
		- cap function returns length of underlying array
		- append functino to add elements to slice
		- may cause extensive copy operation if underlying aray is too small
		-* copies refer to same underlying array
	*/
}
