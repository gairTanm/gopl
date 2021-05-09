package htmlHelpers

import (
	"fmt"

	"golang.org/x/net/html"
)

func PrintText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	if n.FirstChild != nil {
		PrintText(n.FirstChild)
	}
	if n.NextSibling != nil {
		PrintText(n.NextSibling)
	}
}
