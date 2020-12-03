package crawlerutil

import (
	"log"
	"path"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
)

type URLHandler func(url string)

func NewHandlerWithFilters(handler URLHandler, filters ...URLFilter) URLHandler {
	return func(url string) {
		if !filterURL(url, filters...) {
			return
		}

		handler(url)
	}
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
