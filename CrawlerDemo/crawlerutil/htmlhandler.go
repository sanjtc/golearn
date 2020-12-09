package crawlerutil

import (
	"log"
	"net/url"
	"path"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"golang.org/x/net/html"
)

func HandleHref(node *html.Node, c xcrawler.Crawler, etcdInteractor etcd.Interactor) {
	href := htmlutil.GetElementAttributeValue(node, "href")
	key := generateKeyForCrawler(href)

	if !Synchronize(key, etcdInteractor) {
		return
	}

	log.Println("process:", href)

	url, err := url.Parse(href)
	if err != nil {
		log.Println("error:", err)
		return
	}

	hrefHandlers := []HrefHandler{
		NewHandlerWithFilters(handleHrefWithAbsolutePath, filterHrefWithAbsolutePath),
		NewHandlerWithFilters(handleHrefWithRelativePath, filterHrefWithRelativePath),
		NewHandlerWithFilters(handleHrefWithJS, filterHrefWithJS),
	}

	for _, handler := range hrefHandlers {
		if handler(url, c) {
			tryRedirectURL(url)
			downloadURL(url, etcdInteractor)
		}
	}
}

func generateKeyForCrawler(url string) string {
	return path.Join("/crawler", url)
}

func downloadURL(u *url.URL, etcdInteractor etcd.Interactor) {
	if u.String() == "" {
		return
	}

	key := generateKeyForCrawler(path.Join("/download", u.String()))
	if !Synchronize(key, etcdInteractor) {
		return
	}

	if err := htmlutil.DownloadURL(u, path.Join(pathutils.GetModulePath("CrawlerDemo"), "download")); err != nil {
		log.Println("error:", err)
		return
	}

	log.Println("download:", u)
}
