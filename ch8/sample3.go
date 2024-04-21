package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const baseUrl = "https://api.github.com/repos/"

type GitHubIssue struct {
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func fetchGitHubIssue(endpoint string, ch chan<- GitHubIssue, wg *sync.WaitGroup) (*GitHubIssue, error) {

	var data GitHubIssue

	defer wg.Done()

	// startTime := time.Now()
	resp, err := http.Get(baseUrl + endpoint)

	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}

	defer resp.Body.Close()

	// Check if the HTTP request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the body of the response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Errorf("Error decoding response body")
		return nil, err
	}

	ch <- data

	// fmt.Printf("End time: %s\n", time.Since(startTime))
	return &data, nil
}

func main() {
	startNow := time.Now()
	arr := []string{
		"PiotrkasS/Strona/issues/1",
		"nideRbiNsI3/yk/issues/1",
		"victoriaschwin/java-ironhack-final-project/issues/1",
		"kwindrem/TailscaleGX/issues/1",
		"PiotrkasS/Strona/issues/1",
		"nideRbiNsI3/yk/issues/1",
		"victoriaschwin/java-ironhack-final-project/issues/1",
		"kwindrem/TailscaleGX/issues/1",
		"PiotrkasS/Strona/issues/1",
		"nideRbiNsI3/yk/issues/1",
		"victoriaschwin/java-ironhack-final-project/issues/1",
		"kwindrem/TailscaleGX/issues/1",
		"PiotrkasS/Strona/issues/1",
		"nideRbiNsI3/yk/issues/1",
		"victoriaschwin/java-ironhack-final-project/issues/1",
		"kwindrem/TailscaleGX/issues/1",
		"PiotrkasS/Strona/issues/1",
		"nideRbiNsI3/yk/issues/1",
		"victoriaschwin/java-ironhack-final-project/issues/1",
		"kwindrem/TailscaleGX/issues/1",
		"PiotrkasS/Strona/issues/1",
		"nideRbiNsI3/yk/issues/1",
		"victoriaschwin/java-ironhack-final-project/issues/1",
		"kwindrem/TailscaleGX/issues/1",
	}

	ch := make(chan GitHubIssue)
	var wg sync.WaitGroup

	for _, endpoint := range arr {
		wg.Add(1)
		go fetchGitHubIssue(endpoint, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("The operation took", time.Since(startNow))
}
