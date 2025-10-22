package basics

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func pointerss() {
	var a int = 42
	// b stores address of a
	var b *int = &a
	// b Prints address of a
	fmt.Println(a, b)
	// b prints value at a
	fmt.Println(*b)
	a = 27
	fmt.Println(a, *b)
	*b = 43
	fmt.Println(a, *b)

	//Pointer Arithmetics
	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	fmt.Printf("%v %p %p\n", x, y, z)
	// c := &x[1] - 4 // not allowed
	// can use a separate package to do this

	var ms *myStruct // initialized to nil by default
	ms = new(myStruct)
	fmt.Println(ms)
	(*ms).foo = 42
	ms.foo = 45 // compiler allows this too
	fmt.Println((*ms).foo)
	/*
		*POINTERS
			-pointer types use an (*) as prefix
			-use the address of operator (&) to get address of variable
			-Dereferencing pointers
				-Dereference a pointer by preceding with an *
				-Complex types are automatically dereferenced
			-Create pointers to objects
				-can use the address of operator & if value type already exits
					- ms := mystruct{foo:42}
					- p := &ms
				-use addressof operator before initializer
					- &myStruct{foo: 42}
				-use the new keyword
					-cant initialize fields at the same time
			-Types with internal pointers
				-all assignment operations in go are opy operations
				-slices and maps contain internal pointers, so copies point to same underlying data
	*/
}
