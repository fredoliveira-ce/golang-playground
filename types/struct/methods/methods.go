package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	surname string
}

func (p person) getCompleteName() string {
	return p.name + " " + p.surname
}

func (p * person) setName(completeName string) {
	names := strings.Split(completeName, " ")
	p.name = names[0]
	p.surname = names[1]
}

func main() {
	p1 := person{"John", "Balboa"}
	fmt.Println(p1.getCompleteName())

	p1.setName("Rosa Sales")
	fmt.Println(p1.getCompleteName())


}
