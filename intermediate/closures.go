package main

import "fmt"

// function which returns a function
func adder() func() int {
	i := 0
	fmt.Println("Previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

func adderMain() func(int) int {
	i := 0
	fmt.Println("Previous value of i:", i)
	return func(x int) int {
		i += x
		fmt.Println("added ", x, "to i")
		return i
	}
}

type counterOperations struct {
	printValue   func() int
	increaseFunc func(int) int
	decreaseFunc func(int) int
}

func counter() counterOperations {
	i := 0
	return counterOperations{
		printValue: func() int {
			fmt.Println("current value:", i)
			return i
		},
		increaseFunc: func(x int) int {
			i += x
			fmt.Println("added ", x, "to i")
			return i
		},
		decreaseFunc: func(x int) int {
			i -= x
			fmt.Println("removed ", x, "from i")
			return i
		},
	}
}

func closuress() {

	//call our main function only once to initialize
	f := adder()
	f2 := adder()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f2())

	f3 := adderMain()

	fmt.Println(f3(5))
	fmt.Println(f3(1))

	i := counter()
	i.printValue()
	i.increaseFunc(4)
	i.increaseFunc(4)
	i.decreaseFunc(4)

}
