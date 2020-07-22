package main

import "fmt"

func main() {
	p1 := Point{8, 3}
	p2 := Point{23, 7}

	fmt.Println(legs(p1, p2)) // we can see private elements in the same package
	fmt.Println(Distance(p1, p2))
}
