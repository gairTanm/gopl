package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	words, images, err := countWordsAndImages(doc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}

func countWordsAndImages(n *html.Node) (words, images int, err error) {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[len(u)-1]
		u = u[:len(u)-1]
		switch n.Type {
		case html.TextNode:
			words += wordCount(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
	}
	return words, images, nil
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}
