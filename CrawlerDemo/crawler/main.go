package main

import (
	"flag"
	"net/http"

	"github.com/pantskun/golearn/CrawlerDemo/crawlerutil"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

func main() {
	var (
		url             string
		useMultiprocess bool
	)

	flag.StringVar(&url, "url", "https://www.ssetech.com.cn", "url")
	flag.BoolVar(&useMultiprocess, "useMultiprocess", false, "use multiprocess")
	flag.Parse()

	var etcdInteractor etcd.Interactor

	if useMultiprocess {
		i, err := etcd.NewInteractor()
		if err != nil {
			xlogutil.Error(err)
			return
		}

		etcdInteractor = i
	} else {
		i, err := etcd.NewInteractorWithEmbed()
		if err != nil {
			xlogutil.Error(err)
			return
		}

		etcdInteractor = i
	}

	handleHTML := func(element xcrawler.HTMLElement) {
		crawlerutil.HandleElementWithURL(element, etcdInteractor)
	}

	handleReq := func(req *http.Request) {
		crawlerutil.HandleRequestWithSync(req, etcdInteractor)
	}

	c := xcrawler.NewCrawler(2)

	c.AddHTMLHandler(handleHTML, crawlerutil.FilterElementWithURL)
	c.AddRequestHandler(handleReq)

	c.Visit(url)
}
