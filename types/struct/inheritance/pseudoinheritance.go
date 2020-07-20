package main

import "fmt"

type car struct {
	name string
}

type ferrari struct {
	car // anonymous field
	value float64
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.value = 100

	fmt.Println(f)
	fmt.Println(f.name, f.value)
}

