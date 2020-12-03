package crawlerutil

import (
	"log"
	"path"

	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
	"golang.org/x/net/html"
)

func HandleHrefMultiprocess(node *html.Node, etcdInteractor etcd.Interactor) {
	url := htmlutil.GetElementAttributeValue(node, "href")
	key := generateKeyFromURL(url)

	if !Synchronize(key, etcdInteractor) {
		return
	}

	log.Println("process key:", key)

	urlHandlers := []HrefHandler{
		NewHandlerWithFilters(handleHrefWithAbsolutePathAndHTMLFile, filterHrefWithAbsolutePath, filterHrefWithHTMLFile),
		NewHandlerWithFilters(handleHrefWithAbsolutePathAndNotFile, filterHrefWithAbsolutePath, filterHrefWithNotFile),
		NewHandlerWithFilters(handleHrefWithRelativePathAndHTMLFile, filterHrefWithRelativePath, filterHrefWithHTMLFile),
		NewHandlerWithFilters(handleHrefWithRelativePathAndNotFile, filterHrefWithRelativePath, filterHrefWithNotFile),
		NewHandlerWithFilters(handleHrefWithJS, filterHrefWithJS),
	}

	for _, handler := range urlHandlers {
		handler(url)
	}
}

func generateKeyFromURL(url string) string {
	return path.Join("/crawler", url)
}
