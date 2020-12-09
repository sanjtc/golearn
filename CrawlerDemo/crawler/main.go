package main

import (
	"flag"
	"log"

	"github.com/pantskun/golearn/CrawlerDemo/crawlerutil"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"golang.org/x/net/html"
)

func main() {
	var (
		url             string
		useMultiprocess bool
	)

	flag.StringVar(&url, "url", "https://www.ssetech.com.cn/", "url")
	flag.BoolVar(&useMultiprocess, "useMultiprocess", false, "use multiprocess")
	flag.Parse()

	var etcdInteractor etcd.Interactor

	if useMultiprocess {
		i, err := etcd.NewInteractor()
		if err != nil {
			log.Println("error:", err)
			return
		}

		etcdInteractor = i
	}

	handleHref := func(node *html.Node, c xcrawler.Crawler) {
		crawlerutil.HandleHref(node, c, etcdInteractor)
	}

	c := xcrawler.NewCrawler()

	c.AddHTMLHandler(handleHref, crawlerutil.FilterANode)
	c.Visit(url)
}
