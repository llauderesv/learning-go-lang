package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute

	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil // success
		}
		log.SetPrefix("wait: ") // suppresses the date and time
		log.SetFlags(0)
		log.Printf("server not responding (%s); retrying...", err)

		fmt.Printf("Debug: %s\n", time.Duration(tries)*time.Second)
		// time.Sleep(time.Duration(tries) * time.Second)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	url := os.Args[1:]
	if err := WaitForServer(url[0]); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}
