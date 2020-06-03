package main

import (
	"strings"
	"fmt"
	"os"
	"io"
	"net/http"
	"bufio"
	"strconv"
	"time"
	"io/ioutil"
)

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
				printLogs()
			case menuExit: 
				fmt.Println("Exiting.")
				os.Exit(0)
			default: 
				fmt.Println("The command is not valid.")
				os.Exit(-1)
		}
	}
}

func showIntroduction() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
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
	sites := readSiteFromFile()

	for i, site := range sites {
		fmt.Println("Testing site: ", i, ":", site)
		testSite(site)
	}
}

func testSite(site string) {
	response, error := http.Get(site)
	
	if error != nil {
		fmt.Println("Error! Details: ", error)
	}

	if(response.StatusCode == 200) {
		fmt.Println(response)
		writeLogs(site, true)
	} else {
		fmt.Println("ERROR: the website has a problem: ", error, response.StatusCode)
		writeLogs(site, false)
	}
}

func readSiteFromFile() [] string {
	var sites [] string

	file, error := os.Open("sites.txt")

	if(error != nil) {
		fmt.Println("Error, details: ", error)
	}

	reader := bufio.NewReader(file)
	
	for {
		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		
		sites = append(sites, line)

		if error == io.EOF {
			break
		}
	}

	file.Close()
	
	return sites
}

func writeLogs(site string, status bool) {
	file, error := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if error != nil {
		fmt.Println("Error - details: ", error)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	fmt.Println(file)
}

func printLogs() {
	file, error := ioutil.ReadFile("logs.txt") // we don't need to clese when we're using ioutil.ReadFile

	if error != nil {
		fmt.Println(file)
	}

	fmt.Println(string(file))
}