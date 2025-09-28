package main

import "fmt"

type Rectangle struct {
	l float64
	w float64
}

// passed as value
func (r Rectangle) Area() float64 {
	return r.l * r.w
}

// passed by reference using pointers
func (r *Rectangle) Scale(factor float64) {
	r.l *= factor
	r.w *= factor
}

func methodss() {
	rect := Rectangle{l: 10, w: 23}
	area := rect.Area()
	fmt.Println("old area", area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("new area", area)
}
