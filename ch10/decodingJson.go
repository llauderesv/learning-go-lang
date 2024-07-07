package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapeURL        string `json:"unescapeURL"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleURL"`
		CacheURL           string `json:"cacheURL"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func check(err error) {
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
}

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	resp, err := http.Get(uri)
	check(err)

	defer resp.Body.Close()

	// Decode the JSON response into our struct type.
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	check(err)

	fmt.Println(gr)
}
