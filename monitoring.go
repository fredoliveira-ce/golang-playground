package main

import "fmt"
import "os"
import "net/http"

func main()  {	

	showIntroduction()
	
	for {
		command := readInput()

		switch command {
			case 1: 
				initMonitoring()
			case 2: 
				fmt.Println("Showing logs")
			case 0: 
				fmt.Println("Exiting")
				os.Exit(0)
			default: 
				fmt.Println("The command is not valid")
				os.Exit(-1)
		}
	}
}

func showIntroduction() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readInput() int {
	var command int
	fmt.Scanf("%d", &command)
	//fmt.Scan(&command)
	fmt.Println("The command chosen was: ", command)
	fmt.Println("The variable address of the command is: ", &command)

	return command
}

func initMonitoring() {
	fmt.Println("Monitoring...")
	sites := [] string { "https://golang.org/doc/", "https://dictionary.cambridge.org/us/", "https://edition.cnn.com/" }

	fmt.Println(sites)

	for i, site := range sites {
		fmt.Println("I'm at position", i, "of my slice and in this position I have the site:", site)
	}

	site := "https://dictionary.cambridge.org/us/"

	response, error := http.Get(site)
	
	if(response.StatusCode == 200) {
		fmt.Println(response)
	} else {
		fmt.Println("The website:", error, "has a problem", response.StatusCode)
	}
	
}