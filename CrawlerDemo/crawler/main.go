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
	var url string

	flag.StringVar(&url, "url", "https://www.ssetech.com.cn/", "url")
	flag.Parse()

	etcdInteractor, err := etcd.NewInteractor()
	if err != nil {
		log.Println(err)
		return
	}

	handleHref := func(node *html.Node) {
		crawlerutil.HandleHrefMultiprocess(node, etcdInteractor)
	}

	c := xcrawler.NewCrawler()
	c.AddHTMLHandler(handleHref, crawlerutil.FilterANode)
	c.Visit(url)
}
