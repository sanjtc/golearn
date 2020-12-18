package crawlerutil

import (
	"net/url"
	"strconv"

	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

func HandleElementWithURL(element xcrawler.HTMLElement, etcdInteractor etcd.Interactor) {
	s := getURLString(element)

	u, err := url.Parse(s)
	if err != nil {
		xlogutil.Error(err)
		return
	}

	key := generateKeyForCrawler(
		"element",
		"depth:"+strconv.Itoa(element.GetRequest().GetDepth()),
		element.GetRequest().GetRawReq().URL.String(),
		u.String(),
	)

	if !etcdInteractor.TxnSync(key) {
		return
	}

	xlogutil.Warning("process:", key)

	if u.Host == "" {
		u.Host = element.GetRequest().GetRawReq().URL.Host
	}

	if u.Scheme == "" {
		u.Scheme = element.GetRequest().GetRawReq().URL.Scheme
	}

	hrefHandlers := []URLHandler{
		NewHandlerWithFilters(handleURLWithHTTP, filterURLWithHTTP),
		NewHandlerWithFilters(handleURLWithJS, filterURLWithJS),
	}

	for _, handler := range hrefHandlers {
		if handler(u) {
			element.GetRequest().Visit(u.String())
		}
	}
}

func generateKeyForCrawler(keys ...string) string {
	res := "/crawler"
	for _, key := range keys {
		res = res + "-" + key
	}

	return res
}

func getURLString(element xcrawler.HTMLElement) string {
	href := element.GetAttr("href")
	src := element.GetAttr("src")

	switch {
	case href != "":
		{
			return href
		}
	case src != "":
		{
			return src
		}
	default:
		{
			return ""
		}
	}
}
