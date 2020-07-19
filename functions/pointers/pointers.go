package main

import "fmt"

func byCoping(n int) {
	n++
}

// pointer is represented by *
func byReferencing(n *int) {
	// * use to dereference
	// get direct access to the value
	*n++
}

func main() {
	n := 1

	byCoping(n)
	fmt.Println(n)

	// & is used to get the reference
	byReferencing(&n)
	fmt.Println(n)
}
