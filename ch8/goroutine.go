package main

import "fmt"

func main() {

	done := make(chan struct{})

	go func() {
		fmt.Println("World")
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Hello")
}
