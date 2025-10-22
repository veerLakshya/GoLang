package intermediate

import "math"

func typess() {
	var a int = 32
	b := int32(a)
	c := float64(b)

	e := math.Pi
	println(e)
	f := int(e)
	println(f, c)

}
