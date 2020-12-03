package crawlerutil

import (
	"log"
	"path"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
)

type HrefHandler func(string)

func NewHandlerWithFilters(handler HrefHandler, filters ...HrefFilter) HrefHandler {
	return func(href string) {
		if !FilterHref(href, filters...) {
			return
		}

		handler(href)
	}
}

func handleHrefWithAbsolutePathAndHTMLFile(href string) {
	log.Println("AbsolutePathAndHTMLFile: ", href)

	if err := htmlutil.DownloadURL(href, path.Join(pathutils.GetModulePath("CrawlerDemo"), "download")); err != nil {
		log.Println(err)
	}
}

func handleHrefWithAbsolutePathAndNotFile(href string) {
	log.Println("AbsolutePathAndNotFile: ", href)
}

func handleHrefWithRelativePathAndHTMLFile(href string) {
	log.Println("RelativePathAndHTMLFile: ", href)
}

func handleHrefWithRelativePathAndNotFile(href string) {
	log.Println("RelativePathAndNotFile: ", href)
}

func handleHrefWithJS(href string) {
	log.Println("JS: ", href)
}
