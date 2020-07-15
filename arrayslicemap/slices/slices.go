package main

import "fmt"
import "reflect"

func main()  {
	a := [5] int{1, 2, 3, 4, 5}
	s := a[1:3] // same memory reference - not included position 3
	fmt.Println(a, s)

	s2 := a[:2] // from zero to one
	fmt.Println(s2)

	var s3 []int
	s3 = append(s3, 1, 2, 3)
	s4 := make([]int, 10)
	copy(s4, s3)
	fmt.Println(s3, s4)

	showNameWithSlice()

	makeSlice()
}

func makeSlice() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // reaching the maximum cap
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // golang double the capacity
	fmt.Println(s, len(s), cap(s))

	// How an array can be referenced by slices
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2)
}

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