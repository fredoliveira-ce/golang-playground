package main

import (
	"math"
)

// A Point represents a coordinate in a cartesian plane
type Point struct {
	x float64
	y float64
}

func legs(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distance method calculate the distance between two points
func Distance(a, b Point) float64 {
	cx, cy := legs(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
