package basics

import (
	"fmt"
)

// enumerated constants
const (
	// value of this iota is scoped to this block only
	a = iota // iota is a special counter
	b = iota
	c = iota
)
const (
	_  = iota // cant be accessed
	a2 = iota
)

const (
	isAdmin = 1 << iota
	isHead
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeSouthAmerica
)

func constantss() {
	const myConst int = 42
	fmt.Printf("%d, %T\n", myConst, myConst)

	// const myConst float64 = math.Sin(1.57) // cant do this

	//const a int = 14
	//const b string = "foo"
	//const c float32 = 3.14
	//const d bool = true

	const e = 32

	// can do var + const for same type of variables

	var roles byte = isAdmin | canSeeFinancials | canSeeAsia
	fmt.Printf("%b", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&isAdmin == isAdmin)

	// Summary of constants
	/*
		Immutable, but can be shadowed
		Replaced by compiler at compile time so value must be calculable at compile time
		Named like Variable
			-PascalCase for exported constants
			-camelCase for internal constats
		Typed constants work line immutable variables
			- can interoperate only with same type
		Untyped constants work like literals
			- can interoperate with similar types

		Enumerated constants
			- special symbol iota allows related constants to be created easily
			- Iota stats at 0 in each const block and increments by one
			- watch out of constant values that match zero values for variables
		Enumerated Expressions
			-* operations that can be determined at compile time are allowed
			-arithmetic
			-bitwise operations
			-bit shifting
	*/
}
