package main

import (
	"fmt"
	"strconv"
)

// string conversion package

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
	fmt.Printf("%T, %v\n", j, j)

	var theHTTPRequest string = "http://www.google.com"
	fmt.Println(theHTTPRequest)

	// Converting variable types
	var x int = 42
	fmt.Printf("%v, %T\n", x, x)

	var y float32
	y = float32(i) // explicitly convert but chances of loosing type
	fmt.Printf("%v, %T\n", y, y)

	var z string
	// z = string(x) // string is just an alias of bytes so it wouldnt work properly until done using string conversion package
	z = strconv.Itoa(x)
	fmt.Printf("%v, %T\n", z, z)

	// *Variable declaration
	// var foo int
	// var foo int = 32
	// foo := 32

	// cant redeclare variable, but can shadow them
	// all variables must be used otherwise compiler error

	// *Visibility
	// lower case first letter for package scope
	// upper case first letter to export
	// no private scope

	// *Naming Conventions
	// Pascal or camelCase - Capatilize acronyms(HTTP, URL)
	// as short as reasonable, longer names for longer lives

	// *Type Conversions
	// DestinationType(variable)
	// use strconv package for strings

}
