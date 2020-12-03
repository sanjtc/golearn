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

	urlHandlers := []URLHandler{
		NewHandlerWithFilters(handleURLWithAbsolutePathAndHTMLFile, filterURLWithAbsolutePath, filterURLWithHTMLFile),
		NewHandlerWithFilters(handleURLWithAbsolutePathAndNotFile, filterURLWithAbsolutePath, filterURLWithNotFile),
		NewHandlerWithFilters(handleURLWithRelativePathAndHTMLFile, filterURLWithRelativePath, filterURLWithHTMLFile),
		NewHandlerWithFilters(handleURLWithRelativePathAndNotFile, filterURLWithRelativePath, filterURLWithNotFile),
		NewHandlerWithFilters(handleURLWithJS, filterURLWithJS),
	}

	for _, handler := range urlHandlers {
		handler(url)
	}
}

func generateKeyFromURL(url string) string {
	return path.Join("/crawler", url)
}
