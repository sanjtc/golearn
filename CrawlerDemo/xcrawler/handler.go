package xcrawler

import "golang.org/x/net/html"

type HTMLNodeHandler func(*html.Node)
