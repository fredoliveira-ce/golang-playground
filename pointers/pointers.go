package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // get pointer address
	*p++ // get the value
	i++

	fmt.Println(p, *p, i, &i)

}
