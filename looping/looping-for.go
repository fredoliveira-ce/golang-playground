package main

import "fmt"

var sites = [] string { "https://dictionary.cambridge.org/us/", "https://dictionary.cambridge.org/us/", "https://dictionary.cambridge.org/us/" }

func useStandardFor() {
	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
	}
}

func useRangeFor() {
	for i, site := range sites {
		fmt.Println("I'm at position", i, "of my slice and in this position I have the site:", site)
	}
}