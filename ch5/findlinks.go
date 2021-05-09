package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m := make(map[string]int, 0)
	for n, link := range countTypes(m, doc) {
		fmt.Println(n, link)
	}
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	if node.FirstChild != nil {
		links = visit(links, node.FirstChild)
	}
	if node.NextSibling != nil {
		links = visit(links, node.NextSibling)
	}
	return links
}

func countTypes(counts map[string]int, node *html.Node) map[string]int {
	if node.Type == html.ElementNode {
		counts[node.Data]++
	}
	if node.FirstChild != nil {
		countTypes(counts, node.FirstChild)
	}
	if node.NextSibling != nil {
		countTypes(counts, node.NextSibling)
	}
	return counts
}
