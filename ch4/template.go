package main

import (
	"log"
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"wool", 17}

	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.Execute(os.Stdout, sweaters); err != nil {
		log.Fatal(err)
	}
}
