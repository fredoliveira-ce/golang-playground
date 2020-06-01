package main

import "fmt"
import "os"
import "net/http"

func main()  {
	var name string = "Fred"
	var version float32 = 1.1
	var age int = 24

	fmt.Println("Hello, Sr.!", name, ", we're at the vertion", version)
	fmt.Println("Your age is", age)
	
	showIntroduction()
	
	for {
		command := readInput()

		switch command {
			case 1: 
				initMonitoring()
			case 2: 
				fmt.Println("Showing logs")
			case 3: 
				name, age := returnNameAndAge()
				fmt.Println("Name is: ", name, " and Age is: ", age)
			case 4: 
				name, _ := returnNameAndAge()
				fmt.Println("Name is: ", name, " and I don't need the age now.")
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
	fmt.Println("3- Show name and age")
	fmt.Println("4- Show only name")
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

func returnNameAndAge() (string, int) {
	name := "Fred"
	age := 30

	return name, age
}

func initMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://dictionary.cambridge.org/us/"
	response, error := http.Get(site)
	
	if(response.StatusCode == 200) {
		fmt.Println(response)
	} else {
		fmt.Println("The website:", error, "has a problem", response.StatusCode)
	}
	
}