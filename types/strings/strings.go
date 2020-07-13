package main

import "fmt"

func main() {
	s1 := "Hello!"
	fmt.Println("The string size is:", len(s1))

	//backtick
	var strbs = `
		Test
		string 
		with
		backstick
	`
	fmt.Println("The string size is:", len(strbs))

}
