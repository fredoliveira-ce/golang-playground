package main

import (
	"fmt"
	"time"
)

func speak(person, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteraction %d)\n", person, text, i + 1)
	}
}

func main()  {
	go speak("John", "Hi!", 500)
	go speak("Mary", "Hello!", 500)

	time.Sleep(1 * time.Minute) // go routines runs only while the application is running.
}
