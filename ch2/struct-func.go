package main

import "fmt"

type Rectangle struct {
	Width  float32
	Height float32
}

func calculateArea(rect Rectangle) float32 {
	return rect.Width * rect.Height
}

func calculateRadius(diameter int) int {
	return diameter / 2
}

func main() {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	area := calculateArea(rect)

	fmt.Println("Area:", area)

	diameterOfCircumference := calculateRadius(10)
	fmt.Printf("Diameter of Circumference is %d", diameterOfCircumference)

}
