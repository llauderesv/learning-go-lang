package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.333.3333",
	}

	data, err := json.MarshalIndent(c, "", "   ")
	if err != nil {
		log.Panicln(err)
		return
	}

	fmt.Println(string(data))
}
