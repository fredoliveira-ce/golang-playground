package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Tags []string `json:"tags"`
}

func main()  {
	p1 := product{1, "ruler", 0.3, []string{"stationary", "plastic"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 product
	jsonString := `{"id":1,"name":"ruler","price":0.3,"tags":["stationary","plastic"]}`
	_ = json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags)
}
