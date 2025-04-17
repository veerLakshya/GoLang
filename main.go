package main

import "fmt"

// can only declare here at package level with full declaration
var x float64 = 4

// this will get shadowed by same named local variable
var i int = 0

// can be used to categorize or just avoid a lot ussage of var
var (
	actorName    string = "Laksh"
	cmopanion    string = "Badal"
	doctorNumber int    = 3
	season       int    = 1
)

// lower case variable only for this package
// upper case named variable declared at package level can be accessed by anyone consuming that package
// block scope - anything declared between block like main

func main() {
	//Variables in go
	/*
		3 different types -
		1-	var i int
			i = 22
		2-	var j int = 23 		// gives more control
		3 - k := 24				// less control but identifies type
	*/
	var j int = 24
	fmt.Println(i) // gets i from global scope
	i := 2         // make go figure the variable type for us
	fmt.Println(i) // gets from local scope
	fmt.Printf("%T, %v", j, j)
}
