package main

import (
	"flag"
	"log"
	"path"
	"strings"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/crawlerutil"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
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
		handleHrefMultiprocess(node, etcdInteractor)
	}

	c := xcrawler.NewCrawler()
	c.AddHTMLNodeHandler(handleHref, filterANode)
	c.Visit(url)
}

func handleHrefMultiprocess(node *html.Node, etcdInteractor etcd.Interactor) {
	url := htmlutil.GetElementAttributeValue(node, "href")
	key := generateKeyFromURL(url)

	if !crawlerutil.Synchronize(key, etcdInteractor) {
		return
	}

	log.Println("process key:", key)

	urlHandlers := []crawlerutil.URLHandler{
		crawlerutil.NewHandlerWithFilters(handleURLWithAbsolutePathAndHTMLFile, filterURLWithAbsolutePath, filterURLWithHTMLFile),
		crawlerutil.NewHandlerWithFilters(handleURLWithAbsolutePathAndNotFile, filterURLWithAbsolutePath, filterURLWithNotFile),
		crawlerutil.NewHandlerWithFilters(handleURLWithRelativePathAndHTMLFile, filterURLWithRelativePath, filterURLWithHTMLFile),
		crawlerutil.NewHandlerWithFilters(handleURLWithRelativePathAndNotFile, filterURLWithRelativePath, filterURLWithNotFile),
		crawlerutil.NewHandlerWithFilters(handleURLWithJS, filterURLWithJS),
	}

	for _, handler := range urlHandlers {
		handler(url)
	}
}

func filterANode(node *html.Node) bool {
	return node.Data == "a"
}

func generateKeyFromURL(url string) string {
	return path.Join("/crawler", url)
}

func filterURLWithAbsolutePath(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func filterURLWithRelativePath(url string) bool {
	return strings.HasPrefix(url, "/")
}

func filterURLWithHTMLFile(url string) bool {
	ext := path.Ext(url)
	return ext == ".html"
}

func filterURLWithNotFile(url string) bool {
	return strings.HasSuffix(url, "/")
}

func filterURLWithJS(url string) bool {
	return strings.HasPrefix(url, "javascript:")
}

func handleURLWithAbsolutePathAndHTMLFile(url string) {
	log.Println("AbsolutePathAndHTMLFile: ", url)

	if err := htmlutil.DownloadURL(url, path.Join(pathutils.GetModulePath("CrawlerDemo"), "download")); err != nil {
		log.Println(err)
	}
}

func handleURLWithAbsolutePathAndNotFile(url string) {
	log.Println("AbsolutePathAndNotFile: ", url)
}

func handleURLWithRelativePathAndHTMLFile(url string) {
	log.Println("RelativePathAndHTMLFile: ", url)
}

func handleURLWithRelativePathAndNotFile(url string) {
	log.Println("RelativePathAndNotFile: ", url)
}

func handleURLWithJS(url string) {
	log.Println("JS: ", url)
}
