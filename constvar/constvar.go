package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var radius = 3.2
	area := PI * math.Pow(radius, 2)
	fmt.Println("The circumference area is:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 4
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	g, h, i := 2, false, "wow!"

	fmt.Println(e, f, g, h, i)
}
