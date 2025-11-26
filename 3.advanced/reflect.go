package advanced

import (
	"fmt"
	"reflect"
)

func reflection() {
	x := 42
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Value: ", v)
	fmt.Println("Type: ", t)
	fmt.Println("Kind: ", t.Kind() == reflect.Int)
	fmt.Println("Is Zero: ", v.IsZero())

	y := 10
	v1 := reflect.ValueOf(&y).Elem()
	v2 := reflect.ValueOf(y)
	fmt.Println("V2 Type: ", v2.Type())

	fmt.Println("Original value: ", v1.Int())

}

/*
Reflection is a mechanism that allows a program to inspect and manipulate its own structure and behaviour at run time

# Why use reflection?
	- dynamic type inspection
	- generic programming
	- serialization/deserialization

# Few Methods
	- reflect.TypeOf
	- reflect.Value
	- reflect.ValueOf
	- reflect.ValueOf().Elem()
*/
