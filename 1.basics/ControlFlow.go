package basics

import (
	"fmt"
	"math"
)

func floww() {
	// always have to use curly braces with if statements
	if true {
		fmt.Println("this is true")
	}

	statePop := map[string]int{
		"animal":  1,
		"animals": 2,
	}

	// initializer statement, defined only within the scope of this statement
	if pop, ok := statePop["animal"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 30

	// chained if, if else & else, only one code path is evaluated
	if guess < number {
		fmt.Println("too low")
	} else if guess > number {
		fmt.Println("too high")
	} else {
		fmt.Println("too low")
	}

	// be careful with floating points
	num := 0.12323423
	if math.Abs(num/math.Pow(math.Sqrt(num), 2)-1) < 0.001 {
		fmt.Println("these are the same")
	} else {
		fmt.Println("these are the different")
	}

	// * Switch statements
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("four")
	}

	// can have multiple cases but not overlapping
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("three seven or nine")
	}

	// initializer
	switch i := 3 + 2; i {
	case 1, 2:
		fmt.Println("one")
	case 3, 4:
		fmt.Println("two")
	}

	// Tagless syntax, can overlap but only one statement executes as break statement is inbuilt
	i := 10
	switch {
	case i <= 10:
		fmt.Println("i is less than 10")
		fallthrough // allows cases ahead to execute instead of breaking here by default
	case i <= 20:
		fmt.Println("i is less than 20")
	case i <= 30:
		fmt.Println("i is less than 30")
	default:
		fmt.Println("i is greater than 30")
	}

	// checking type of interface and executing accordingly
	var x interface{} = 1
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
		//break, use if we want to stop at a specific point in a case
		fmt.Printf("this will print too if break is not used above")
	case float64:
		fmt.Println("x is a float")
	default:
		fmt.Println("x is another type")
	}

	/*
		*IF statements*
			-initializer
			-comparison operators
			-logical operators
			-short-circuiting
			-if-else statements
			-if-else if statements
			-equality and floats

		*SWITCH Statements*
			-switching on a tag
			-cases with multiple tests
			-initializers
			-switches with no tag
			-implicit break is inbuilt
			-fallthrough keyword
			-type switches
			-breaking out early within a case using break
	*/

}
