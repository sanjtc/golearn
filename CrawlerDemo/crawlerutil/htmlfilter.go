package crawlerutil

import "golang.org/x/net/html"

func FilterANode(node *html.Node) bool {
	return node.Data == "a"
}
