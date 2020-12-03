package xcrawler

import "golang.org/x/net/html"

type HTMLNodeFilter func(*html.Node) bool

func FilterHTMLNode(node *html.Node, filters ...HTMLNodeFilter) bool {
	for _, filter := range filters {
		if !filter(node) {
			return false
		}
	}

	return true
}
