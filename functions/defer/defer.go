package main

import "fmt"

func getValue(number int) int {
	defer fmt.Println("end!")
	if number > 1000 {
		fmt.Println("High")
		return 5000
	}
	fmt.Println("Low")
	return number
}

func main() {
	fmt.Println(getValue(2000))
	//fmt.Println(getValue(1000))
}
