package main

import "fmt"

func average(numbers ...float64) float64  {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total/float64(len(numbers))
}

 func main() {
	fmt.Printf("Average: %.2f", average(7.7, 6.3, 6.7, 9.5))

	sliceNumbers := []float64{23.2, 23.1, 45.2, 45.2}
	 fmt.Printf("Average: %.2f", average(sliceNumbers...))
}
