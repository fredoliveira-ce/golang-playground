package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

func main() {
	useGetRequest()
	useGetRequestWithHeader()
}

type Person struct {
	id string `json:"id"`
	nome string `json:"nome"`
	endereco  string `json:"endereco"`
}

func useGetRequest() {
	site := "https://dictionary.cambridge.org/us/"
	response, error := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println(response)
	} else {
		fmt.Println("The website:", error, "has a problem", response.StatusCode)
	}
}

func useGetRequestWithHeader() {
	auth := "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiI5MjQ1MjY2MzM4NyIsImV4cCI6MTYxMzI3MDU2NX0.Y8qMxdab2BvPxK-2W0Bn3Ev_A1eaWVfJ_-yS5eof0HoGQp8Qb9Z41LyIk6QJB3xEL5N0h6KBwxAZloLroG2zQQ"
	url := "https://servicos.sda.ce.gov.br/dataservices/api/pessoa/75850664300"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", auth)
	res, _ := client.Do(req)
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
