package basics

import "fmt"

func loops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	//while loops in go
	for i < 7 {
		fmt.Println(i)
		i++
	}
	//another way
	for {
		i++
		if i == 9 {
			continue
		}
		fmt.Println(i)

		if i > 10 {
			break
		}
	}
	//Nested Loops
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}
	s := []int{1, 2, 3}
	// for range loop based on key, value pair
	for _, v := range s {
		fmt.Println(v)
	}

	/*
		Simple loops
			-for initializer, test, incrementer {}
			-for test {}
			-for {}
			-Exiting Early
				-break
				-continue
				-labels
		Looping over collections
			-arrays, slices, maps, strings, channels
			-for k, v := range collection {}
	*/
}
