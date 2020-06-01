package main

import "fmt"

func showFixedArray() {
	var sites [4] string
	sites[0] = "https://dictionary.cambridge.org/us/"
	sites[1] = "https://edition.cnn.com/"
	sites[2] = "https://golang.org/doc/"
	sites[3] = "https://editor.swagger.io/"

	fmt.Println(sites)
}