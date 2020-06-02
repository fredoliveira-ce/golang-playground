package main

import "fmt"
import "os"
import "net/http"

const menuInitMonitoring = 1
const menuShowLogs = 2
const menuExit = 0

func main()  {	
	for {
		showIntroduction()

		command := readInput()

		switch command {
			case menuInitMonitoring: 
				initMonitoring()
			case menuShowLogs: 
				fmt.Println("Showing logs")
			case menuExit: 
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
	fmt.Println("The command chosen was: ", command)
	fmt.Println("The variable address of the command is: ", &command)

	return command
}

func initMonitoring() {
	fmt.Println("Monitoring...")
	sites := [] string { "https://golang.org/doc/", "https://dictionary.cambridge.org/us/", "https://edition.cnn.com/" }

	fmt.Println(sites)

	for i, site := range sites {
		fmt.Println("Testing site: ", i, ":", site)
		testSite(site)
	}
}

func testSite(site string) {
	response, error := http.Get(site)
	
	if(response.StatusCode == 200) {
		fmt.Println(response)
	} else {
		fmt.Println("ERROR: the website: ", error, "has a problem", response.StatusCode)
	}
}