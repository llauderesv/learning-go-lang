package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // sending values to the channel
		}

		close(ch)
	}()

	for num := range ch {
		fmt.Println("Received: ", num)
	}
}
