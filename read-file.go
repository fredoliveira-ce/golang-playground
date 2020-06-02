package main

import "fmt"
import "io/ioutil"
import "os"
import "bufio"

func readLineByLine() {
	file, error := os.Open("sites.txt")

	if(error != nil) {
		fmt.Println("Error, details: ", error)
	}

	reader := bufio.NewReader(file)
	line, error := reader.ReadString('\n')

	if error != nil {
		fmt.Println("Error, details: ", error)
	}

	fmt.Println(line)
}

func readAllFile() {
	file, error := ioutil.ReadFile("sites.txt")

	if error != nil {
		fmt.Println("Error, details: ", error)
	}

	fmt.Println(file)
}

