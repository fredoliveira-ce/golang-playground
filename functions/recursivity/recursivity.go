package main

import "fmt"

func factorial(n uint) uint  {
	switch {
	case n == 0:
		return 1
	default:
		return n * factorial(n - 1)
	}
}

func main() {
	result := factorial(5)
	fmt.Println(result)
}
