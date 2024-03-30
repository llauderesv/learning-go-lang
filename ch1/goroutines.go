package main

import "fmt"

func square(num int, ch chan int) {
	result := num * num
	ch <- result
}

func main() {
	ch := make(chan int)
	go square(3, ch)

	result := <-ch
	fmt.Println(result)
}
