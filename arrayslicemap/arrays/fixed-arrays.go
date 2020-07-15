package main

import "fmt"

func main() {
	numbers := [...] int{1, 2, 3, 4, 5} // different from slices - size will be decided at compilation time

	for i, number := range numbers {
		fmt.Printf("%d) %d\n", i, number)
	}

	for _, number := range numbers {
		fmt.Println(number)
	}

	for index := range numbers {
		fmt.Println(index)
	}
}

func showFixedArray() {
	var sites [4] string
	sites[0] = "https://dictionary.cambridge.org/us/"
	sites[1] = "https://edition.cnn.com/"
	sites[2] = "https://golang.org/doc/"
	sites[3] = "https://editor.swagger.io/"

	fmt.Println(sites)
}