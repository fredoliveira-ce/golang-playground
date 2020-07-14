package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(getScore(10))
	fmt.Println(getScore(4))

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}

func getScore(n float64) string {
	var note = int(n)
	switch note {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 4, 3:
		fallthrough
	case 2, 1:
		return "C"
	default:
		return "Invalid note"
	}
}
