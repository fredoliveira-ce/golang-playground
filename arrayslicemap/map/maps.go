package main

import "fmt"

func main() {
	folks := make(map[int]string)

	folks[1] = "John"
	folks[2] = "Jessica"
	folks[3] = "Ana"

	for id, name := range folks {
		fmt.Printf("id: %d, name: %s\n", id, name)
	}

	delete(folks, 1)
	delete(folks, 11345) // no problem
	fmt.Println(folks)

	months := map[int]string {
		1: "JAN",
		2: "FEB",
		3: "MAR",
	}

	months[4] = "APR"

	nestedMap()
}

func nestedMap() {
	semesters := map[int]map[int]string {
		1: {
			1: "JAN",
			2: "FEB",
		},
		2: {
			3: "MAR",
			4: "APR",
		},
	}

	fmt.Println(semesters)
}
