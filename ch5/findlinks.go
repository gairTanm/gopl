package main

import (
	"fmt"
	htmlHelpers "gopl/ch5/html"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// m := make(map[string]int, 0)
	for n, link := range htmlHelpers.Visit(nil, doc) {
		fmt.Println(n, link)
	}
}
