package main

func main() {

	// default value is ptr is nil
	var ptr *int
	println(ptr)
	// &a gives the address of variable named a			// referencing
	var a int = 1
	ptr = &a
	// *prt gives the value stored at address ptr		// Dereferencing
	var val int = *ptr
	println(val)

	modifyValue(ptr)
	println(a)
}

func modifyValue(ptr *int) {
	*ptr += 5
}
