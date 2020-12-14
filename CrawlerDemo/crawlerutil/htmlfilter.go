package crawlerutil

import (
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
)

func FilterElementWithURL(element xcrawler.HTMLElement) bool {
	return element.GetAttr("href") != "" || element.GetAttr("src") != ""
}
