package main

import "fmt"

type Duck interface {
	Quack()
}

type Donald struct {
	Name string
}

func (d *Donald) Quack() {
	fmt.Printf("My Name is: %s Quack! Quack!", d.Name)
}

func main() {
	d := Donald{Name: "Vince Llauderes"}

	runDuck(&d)
}

func runDuck(d Duck) {
	d.Quack()
}
