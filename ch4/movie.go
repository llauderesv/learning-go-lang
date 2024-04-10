package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

/*
Converting a Go data structure like movies to JSON is called marshaling.

*/

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "In-grid Bergman"}},
		{Title: "Young Sheldong", Year: 2001, Color: true, Actors: []string{"Annie Potts", "Sheldon Cooper"}},
	}

	// data, err := json.MarshalIndent(movies, "", "") // Format the JSON with identation
	data, err := json.Marshal(movies) // Format the JSON with identation
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)

}
