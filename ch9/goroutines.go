package main

import (
	"fmt"
	"sync"
	"time"
)

// Pass the waitGroup as pointer
func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// Sleep the main thread
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// Wait group is code in Golang to use to wait the execution of go routines.
	// wg.Add telling that how much wait group that we need to wait
	var wg sync.WaitGroup
	wg.Add(1)
	go say("world", &wg)
	wg.Wait()
}
