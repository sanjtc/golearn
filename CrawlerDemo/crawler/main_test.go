package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

func TestURL(t *testing.T) {
	s := "http://www.test.com"

	u, _ := url.Parse(s)

	xlogutil.Warning(u)
}

func TestRedirect(t *testing.T) {
	resp, _ := http.Get("https://www.ssetech.com.cn/place/redirect.do?id=57")
	t.Log(resp.Request.URL)
}
