package xcrawler

import "golang.org/x/net/html"

type HTMLFilter func(*html.Node) bool

func FilterHTML(node *html.Node, filters ...HTMLFilter) bool {
	for _, filter := range filters {
		if !filter(node) {
			return false
		}
	}

	return true
}
