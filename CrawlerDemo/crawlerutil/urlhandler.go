package crawlerutil

import (
	"net/url"

	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

type URLHandler func(u *url.URL) bool

func NewHandlerWithFilters(handler URLHandler, filters ...URLFilter) URLHandler {
	return func(u *url.URL) bool {
		if !FilterURL(u, filters...) {
			return false
		}

		return handler(u)
	}
}

func handleURLWithHTTP(u *url.URL) bool {
	xlogutil.Warning("HandleHrefWithHTTP:", u.String())

	return true
}

func handleURLWithJS(u *url.URL) bool {
	xlogutil.Warning("HandleHrefWithJS:", u.String())

	return false
}
