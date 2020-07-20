package main

import "fmt"

type score float64

func (s score) between(from, to float64) bool {
	return float64(s) >= from && float64(s) <= to
}

func where(s score) string {
	if s.between(1, 10) {
		return "You got it!"
	} else {
		return "Try again!"
	}
}

func main()  {
	fmt.Println(where(0))
	fmt.Println(where(3))
}
