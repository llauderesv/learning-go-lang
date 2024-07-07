package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address struct {
		ZipCode string `json:"zipCode"`
		Street  string `json:"street"`
		City    string `json:"city"`
	} `json:"address"`
}

var JSON = `{
	"name": "Vincent Llauderes",
	"age": "17",
	"address": {
		"zipCode": "1400",
		"street": "General Malvar",
		"city": "Caloocan City"
	}
}`

func main() {
	var p Person
	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("OUTPUT: %v", p)
}
