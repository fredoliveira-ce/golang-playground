package main

import "fmt"

func printNameAndAge() {
	name, age := returnNameAndAge()
	fmt.Println("Name is: ", name, " and Age is: ", age)
}

func printOlnyName() {
	name, _ := returnNameAndAge()
	fmt.Println("Name is: ", name, " and I don't need the age now.")
}
			
func returnNameAndAge() (string, int) {
	name := "Fred"
	age := 30

	return name, age
}