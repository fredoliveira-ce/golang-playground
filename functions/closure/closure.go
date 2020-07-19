package main

import "fmt"

func closure() func() {
	x:=10
	var function = func() {
		fmt.Println(x)
	}
	return function
}

func main() {
	x := 20
	fmt.Println(x)

	printX := closure()
	printX()
}
