package xcrawler

import "golang.org/x/net/html"

type HTMLHandler func(*html.Node, Crawler)
