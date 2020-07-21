package main

import "fmt"

type machine interface {
	turnOn() string
}

type eletronic interface {
	plugIn() string
}

type eletronicMachine interface {
	machine
	eletronic
}

type blender struct {
	brand string
}

type coffeepot struct {
	color string
}

func (b blender) turnOn() string {
	return "Blender is crushing a fruit"
}

func (b blender) plugIn() string {
	return "blender plugged"
}

func (c coffeepot) turnOn() string {
	return "Coffeepot is preparing a cup of coffee"
}

func (b coffeepot) plugIn() string {
	return "coffeepot plugged"
}

func operate(m eletronicMachine) {
	fmt.Println(m.turnOn())
}

func main() {
	operate(blender{})
	operate(coffeepot{})

	var b eletronicMachine = blender{brand: "Oster"}
	fmt.Println(b.turnOn())

	var c machine = coffeepot{color: "Red"}
	fmt.Println(c.turnOn())
}
