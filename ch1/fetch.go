// Fetch prints the contents found at a URL.

package main

import (
	"fmt"
	"io"
	"strings"

	// "io/ioutil"
	"net/http"
	"os"
)

// exercise 1.8
func completeUrl(str string) string {
	if strings.HasPrefix(str, "http://") {
		return str
	}
	return fmt.Sprintf("http://%s", str)
}

func main() {
	for _, url := range os.Args[1:] {

		resp, err := http.Get(completeUrl(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// exercise 1.7
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		// exercise 1.9
		fmt.Println("Status code: %s", resp.Status)
		fmt.Printf("%s", b)
	}
}
