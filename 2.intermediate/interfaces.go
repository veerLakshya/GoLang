package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	h, w float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.h * r.w
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rect) perimeter() float64 {
	return 2 * (r.h + r.w)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) diameter() float64 {
	return 2 * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func interfacess() {
	r := rect{w: 2, h: 3}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
