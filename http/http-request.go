package http

import "fmt"
import "net/http"

func useGetRequest() {
	site := "https://dictionary.cambridge.org/us/"

	response, error := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println(response)
	} else {
		fmt.Println("The website:", error, "has a problem", response.StatusCode)
	}
}