package main

import (
	"net/url"
	"testing"

	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

func TestURL(t *testing.T) {
	s := "http://www.test.com"

	u, _ := url.Parse(s)

	xlogutil.Warning(u)
}
