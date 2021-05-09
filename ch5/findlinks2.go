package main

import (
	"fmt"
	htmlHelpers "gopl/ch5/html"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Server status %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return htmlHelpers.Visit(nil, doc), nil
}
