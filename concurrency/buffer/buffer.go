package main

import "fmt"

func routine(ch chan int) {
	ch <- 1
	fmt.Println("Executed 1")
	ch <- 2
	fmt.Println("Executed 2")
	ch <- 3
	fmt.Println("Executed 3")
	ch <- 4
	fmt.Println("Blocked ")
	ch <- 5
	fmt.Println("Will not execute until consume another one")
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go routine(ch)

	fmt.Println(<-ch)
}
