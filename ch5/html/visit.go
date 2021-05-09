package htmlHelpers

import "golang.org/x/net/html"

func Visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && (node.Data == "a" || node.Data == "img" || node.Data == "script") {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			} else if attr.Key == "src" {
				links = append(links, attr.Val)
			}
		}
	}
	stylesheet := false
	if node.Type == html.ElementNode && (node.Data == "link") {
		for _, attr := range node.Attr {
			if attr.Key == "rel" {
				if attr.Val == "stylesheet" {
					stylesheet = true
				}
			} else if attr.Key == "href" && stylesheet {
				links = append(links, attr.Val)
			}
		}
	}
	if node.FirstChild != nil {
		links = Visit(links, node.FirstChild)
	}
	if node.NextSibling != nil {
		links = Visit(links, node.NextSibling)
	}
	return links
}
