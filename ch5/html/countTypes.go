package htmlHelpers

import "golang.org/x/net/html"

func CountTypes(counts map[string]int, node *html.Node) map[string]int {
	if node.Type == html.ElementNode {
		counts[node.Data]++
	}
	if node.FirstChild != nil {
		CountTypes(counts, node.FirstChild)
	}
	if node.NextSibling != nil {
		CountTypes(counts, node.NextSibling)
	}
	return counts
}
