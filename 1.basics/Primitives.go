package basics

import (
	"fmt"
)

// Boolean type
// Numeric type - int, floating point, complex numbers
// Text Types

func primitives() {
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

	// *Complex Numbers
	// 2 types - complex64, complex 128(float64 + float64)
	var com complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", com, com)
	fmt.Printf("%v, %T\n", real(com), real(com)) // another way to get real and complex part

	// *Strings - any utf-8 character
	s := "Hello Go"
	// s[2] = "u" // cant change strings, IMMUTABLE
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[1], s[1]) // will give int form as strings are alias of bits
	// cant do string(s[2]) to get this

	// cant do string concatenation with +
	// can convert them to collection of bytes
	s1 := []byte(s)
	fmt.Printf("%v, %T\n", s1, s1)

	var r rune = 'a'
	// RUNE in go - represents a single utf-8 char(same as int32)
	fmt.Println(r)        // prints 97
	fmt.Printf("%c\n", r) // prints a as %c used

	// if datastream is in utf-32 then we use special fns that return those values
	// ReadRune from go package

	// Summary-
	/*
			//Booleans
			Boolean type - true or false
			not an alias for other types
			zero value is false

			// Numerics
			Integers
				-Signed Integers
				- int type has varying size, but min 32 bits
				- 8 bit (int8) through 64 bits
			Unsigned integers
				- 8 bit(byte and uint8) through 32bit(uint32)

			cant mix types

		//Text Types
			Strings
				-utf-8
				-immutable
				-can be concatenated with +
				-can be converted to []byte
			Rune
				-utf-32
				-alias for int32
				-special methods are required to process normally

	*/

}
