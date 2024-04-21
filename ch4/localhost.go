package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const apiURL = "https://dummyjson.com/users"

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Position string `json:"position"`
}

type PeopleResponse struct {
	People []*Person
}

func fetchApi(ch chan<- []Person, wg *sync.WaitGroup) ([]Person, error) {

	defer wg.Done()

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	// This defer command will
	defer resp.Body.Close() // Will call after the all the commands below before the return statement.

	// Check if the HTTP request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON data into the struct
	var apiResponse []Person
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	ch <- apiResponse

	return apiResponse, nil
}

func main() {

	start := time.Now()
	ch := make(chan []Person)
	var wg sync.WaitGroup
	// var peopleArray [][]Person

	for i := 0; i <= 10; i++ {
		wg.Add(1) // increments go routine to wait
		go fetchApi(ch, &wg)
		// peopleArray = append(peopleArray, people)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// Simulate modifying the first person's name
		// In a real scenario, ensure the slice is not empty to avoid a panic
		// if len(people) > 0 {
		// 	people[0].Name = "Vince"
		// 	fmt.Printf("Updated first person's name to: %s\n", people[0].Name)
		// } else {
		// 	fmt.Println("No people found in the response")
		// }
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for arr := range ch {
		fmt.Printf("API Response: %v\n", arr)
	}

	fmt.Println("Finished:", time.Since(start))

	// fmt.Printf("API Response: %+v\n", people)
	// For demonstration, print out the modified list
	// for i, person := range people {
	// 	fmt.Printf("Person %d: %s\n", i, person.Name)
	// }
}
