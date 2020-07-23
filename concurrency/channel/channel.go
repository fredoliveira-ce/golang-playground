package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int, 1)

	ch <- 1 // sending data to the channel
	<-ch // reading

	ch <- 2
	fmt.Println(<-ch)

	ch2 := make(chan string)
	go call1(ch2)
	go call2(ch2)

	a, b := <-ch2, <-ch2
	fmt.Println(a, b)
	//fmt.Println(<-ch2) if I'd call again it'll produce an error (deadlock)
}

func call1(c chan string) {
	time.Sleep(time.Second)
	c <- "Call 1"
}

func call2(c chan string) {
	time.Sleep(time.Second)
	c <- "Call 2"
}
