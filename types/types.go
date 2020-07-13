package main

import (
	"fmt"
	"math"
)

func main() {
	// unsigned int (uint)
	/*
	only positives:
	- uint8 (byte - 8 bits)
	- uint16 (short - 2 bytes)
	- uint32 (int - 4 bytes)
	- uint64 (long - 8 bytes)
	*/

	// signed int (int)
	i1 := math.MaxInt64
	fmt.Println("The maximum int value is", i1)

	// initial values

	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("\n%v %v %v %v %q %v", a, b, b, c, d, e)


}
