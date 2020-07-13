package main

import "fmt"

func main() {

	a := 1
	b := 1.9999
	c := false
	d := "ops"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)
}
