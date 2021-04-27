package main

// typical use of fetch, ./fetch <url>
// fetches plaintext view of the website
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error occurred:", err)
			os.Exit(1)
		}
		fmt.Println("Server response status", resp.Status)

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Println("Error occurred while reading the url:", err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
