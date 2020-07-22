package main

import "fmt"

func main()  {
	ch := make(chan int, 1)

	ch <- 1 // sending data to the channel
	<-ch // reading

	ch <- 2
	fmt.Println(<-ch)
}
