package main

import (
	"fmt"
	"strconv"
)

func main() {
	// warning
	fmt.Println("Test" + string(123)) // print Test { (ASC table)

	// int to string
	fmt.Println("Test" + strconv.Itoa(123)) // print Test 123

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	b2, _ := strconv.ParseBool("1")
	if b && b2 {
		fmt.Println(b, b2)
	}
}
