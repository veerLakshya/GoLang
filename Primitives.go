package main

import (
	"fmt"
)

// Boolean type
// Numeric type - int, floating point, complex numbers
// Text Types

func main() {
	// *Booleans
	var f bool = false // if not given value while declaring, intialized with 0
	f = 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", m, m)

	// *Numeric Types
	// int8 -> -128 to 127
	// int16 ->
	// int32 ->
	// for bigger, big package from math library

	var x uint = 42
	fmt.Printf("%v, %T\n", x, x)
	// uint8 -> 0 to 255
	// uint16 -> 0 - 65,535
	// uint32 -> 0 - 4,294,967,295

	// *Arithmetic Operations
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// cant add two different types of integers
	// Ex - int + int8, will have to type conversion

	// *Bit Operators
	fmt.Println(a & b)  // same bits remain on in result
	fmt.Println(a | b)  // bits on in anyone remains on in result
	fmt.Println(a ^ b)  // places with alternate bits remain on in result
	fmt.Println(a &^ b) // places where neither of numbers have bits on remain on in result

	fmt.Println(a << 3) // shift all bits to left by 3
	fmt.Println(a >> 3) // shift all bits to right by 3

	// *Floating Point Numbers
	// float32, float64
	n := 3.14 // always initialized as float64
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)
}
