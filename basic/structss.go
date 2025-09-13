package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"20"` //
	Origin string
}

// Bird not traditional inheritance
type Bird struct {
	Animal   // embedding another struct in
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 0.5,
		CanFly:   true,
	}

	//b.Name = "Emu"
	//b.Origin = "Australia"
	//b.SpeedKPH = 48
	//b.CanFly = true

	fmt.Println(b)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Origin")
	fmt.Println(field.Tag)

	/*
		*STRUCTS*
			-collection of disparate data types that describe a single concept
			-keyed by name fields
			-normally created as types, but anonymous structs are allowed
			- structs are value types
			- no inheritance, but can use composition via embedding
			- tags can be added to struct fields to describe field
	*/
}
