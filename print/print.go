package main

import "fmt"

func main() {

	a := 1
	b := 1.9999
	c := false
	d := "ops"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)
	fmt.Println("")
	fmt.Println(formatAndReturn("test", "return"))
}

func formatAndReturn(p1, p2 string) string {
	return fmt.Sprintf("Return: %s and %s", p1, p2)
}
