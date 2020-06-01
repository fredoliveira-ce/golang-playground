package main

import "fmt"
import "reflect"

func showVariableType()  {
	var name string = "Fred"
	company := "Fred"
	var version float32 = 1.1
	var aFloatValue = 1.1

	fmt.Println("The type of the variable name is: ", reflect.TypeOf(name))
	fmt.Println("The type of the variable aFloatValue is: ", reflect.TypeOf(aFloatValue))
	fmt.Println("The type of the variable company is: ", reflect.TypeOf(company))
	fmt.Println("The type of the variable version is: ", reflect.TypeOf(version))
}