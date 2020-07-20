package main

import "fmt"

type product struct {
	name string
	price float64
	discount float64
}

func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var p product
	p = product {
		name: "pencil",
		price: 1.80,
		discount: 0.05,
	}
	fmt.Println(p, p.priceWithDiscount())

	p2 := product{"notebook", 10.9, 0.10}
	fmt.Println(p2.priceWithDiscount())
}
