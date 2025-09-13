package main

import (
	"errors"
	"fmt"
)

func sayMsg(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(msg)
	}
}

// Variadic parameter, can have only 1 and at last in function
func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

// entry point, takes no parameter, returns nothing
func main() {
	sayMsg("Hello go", 5)
	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)

	fmt.Println(divide(1, 2))

	// immediately used anonymous function
	func() {
		fmt.Println("Hello go")
	}()

	// declaring then initializing function
	var div func(float64, float64) (float64, error)
	div = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("divide by zero")

		} else {
			return a / b, nil
		}
	}

	d, err := div(1, 2)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(d)

	/*
		FUNCTIONS
			-func foo(){
			...
			}
		*Parameters
			-comma-delimited list of variables and types
				-func foo(bar string, baz int)
			-parameters of same type list type once
				-func foo(bar, baz int)
			-when pointers are passed in, the function can change the value in the caller
				-this is always true for maps and slices
			-use variadic parameters to send list of same values in
				-must be last parameter
				-received as a slice
				-func foo(bar string, baz int...)

		*Return Values
			-single return values just list type
				-func foo() int
			-multiple return value list types surrounded by parentheses
				-func foo() (int, error)
				- the (result type, error) paradigm is a very common idiom
			-can use named return values
				-initializes returned values
				-return using return keyword on its own
			-can return addresses of local variables
				-automatically promoted from local memory(stack) to shared memory(heap)

		*Anonymous functions
			-functions don't have names if they are
				-invoked immediately
				-assigned to a variable or passed as argument to a function

		*Functions as types
			-can assign functions to variables or use as arguments and return values in functions
			-type signature is like function signature, with no parameter names
				- var f func(string, string, int)(int, error)

		*Methods
		-functions that execute in context of a type
		-format
			-func(g greeter) greet(){...}
		-receiver can be value or pointer
			-value receiver gets copy of type
			-pointer receiver gets pointer of type

	*/
}
