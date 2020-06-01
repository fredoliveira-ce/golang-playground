package main

import "fmt"
import "reflect"

func showNameWithSlice() {
	names := [] string { "John", "Daniel", "Mathews" }

	fmt.Println(names)
	fmt.Println(reflect.TypeOf(names))
	fmt.Println("The slice size is", len(names))
	fmt.Println("The slice has capability to hold", cap(names), "items")

	names = append(names, "Jessica")

	fmt.Println(names)
	fmt.Println(reflect.TypeOf(names))
	fmt.Println("The slice size is", len(names))
	fmt.Println("The slice has capability to hold", cap(names), "items")
}