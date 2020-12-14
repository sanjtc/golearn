package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
	"github.com/pantskun/golearn/CrawlerDemo/xhttputil"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
	"github.com/robertkrimen/otto"
	"golang.org/x/net/html"
)

func TestURL(t *testing.T) {
	s := "http://www.test.com"

	u, _ := url.Parse(s)

	xlogutil.Warning(u)
}

func TestMP4(t *testing.T) {
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			t.Log("req: ", req)
			t.Log("via num: ", len(via))
			t.Log("via: ", via[0])
			return http.ErrUseLastResponse
		},
	}

	u, _ := url.Parse("https://www.ssetech.com.cn/statics/upload/2019/05-31/16-28-140618599986926.mp4")
	req := &http.Request{
		Method: "GET",
		URL:    u,
	}

	resp, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if err := htmlutil.WriteToFile("test.mp4", body); err != nil {
		t.Fatal(err)
	}
}

func TestNoCookie(t *testing.T) {
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			t.Log("req: ", req)
			t.Log("via num: ", len(via))
			t.Log("via: ", via[0])
			return http.ErrUseLastResponse
		},
	}

	// 1
	u, _ := url.Parse("https://beian.miit.gov.cn/")
	req := &http.Request{
		Method: "GET",
		URL:    u,
		Header: make(http.Header),
	}

	resp1, err := c.Do(req)
	if err != nil {
		xlogutil.Error(err)
		return
	}

	defer resp1.Body.Close()

	var (
		cookie1 *http.Cookie
		cookie2 *http.Cookie
	)

	body1, _ := ioutil.ReadAll(resp1.Body)

	nodes := GetScriptNodes(string(body1))

	for _, node := range nodes {
		c := GetCookieFromScriptNode(node)
		if c != nil {
			cookie1 = c
			break
		}
	}

	setCookie := resp1.Header.Get("Set-Cookie")
	cookie2 = xhttputil.ParseCookie(setCookie)

	req.AddCookie(cookie1)
	req.AddCookie(cookie2)

	// 2
	resp2, err := c.Do(req)
	if err != nil {
		t.Log(err)
		return
	}
	defer resp2.Body.Close()

	body2, _ := ioutil.ReadAll(resp2.Body)

	nodes = GetScriptNodes(string(body2))
	s := GetScriptFromScriptNode(nodes[0])

	htmlutil.WriteToFile("originJS2.js", []byte(s))

	script := ReplaceJS(s)

	htmlutil.WriteToFile("replaceJS2.js", []byte(script))

	v := GetVarValueFromJS(script, "wanted")
	t.Log(v)
}

func GetScriptNodes(body string) []*html.Node {
	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		xlogutil.Error(err)
		return nil
	}

	doc := goquery.NewDocumentFromNode(root)

	s := doc.Find("Script")
	if s == nil {
		return nil
	}

	return s.Nodes
}

func GetCookieFromScriptNode(node *html.Node) *http.Cookie {
	contentNode := node.FirstChild

	if contentNode == nil {
		return nil
	}

	content := contentNode.Data

	newContent := strings.Replace(content, "document.cookie", "cookie", 1)

	s, err := GetVarValueFromJS(newContent, "cookie").ToString()
	if err != nil {
		return nil
	}

	return xhttputil.ParseCookie(s)
}

func GetScriptFromScriptNode(node *html.Node) string {
	contentNode := node.FirstChild

	if contentNode == nil {
		return ""
	}

	return contentNode.Data
}

func GetVarValueFromJS(js string, varname string) otto.Value {
	vm := otto.New()
	vm.Run(
		js,
	)

	value, err := vm.Get(varname)
	if err != nil {
		return otto.NullValue()
	}

	return value
}

func TestJS(t *testing.T) {
	jsfile, err := os.Open("test.js")
	if err != nil {
		t.Log(err)
		return
	}
	defer jsfile.Close()

	jsByte, _ := ioutil.ReadAll(jsfile)
	jsContent := string(jsByte)

	// 替换document[]
	jsContentAfterReplaceDocument := jsContent

	documentIndex := strings.Index(jsContent, "document")
	if documentIndex != -1 {
		equalAfterDocumentIndex := strings.Index(jsContent[documentIndex:], "=")
		if equalAfterDocumentIndex != -1 {
			replaceOld := jsContent[documentIndex : documentIndex+equalAfterDocumentIndex]
			t.Log(replaceOld)
			jsContentAfterReplaceDocument = strings.Replace(jsContent, replaceOld, "wanted", 1)
		}
	}

	// 替换setTimeout
	justDoIt := "function just_do_it(p1, p2) { p1(); }"
	finalJS := strings.Replace(jsContentAfterReplaceDocument, "setTimeout", "just_do_it", 1)

	script := "var wanted;" + justDoIt + finalJS

	// _ = htmlutil.WriteToFile("replaceJS.js", []byte(script))

	value := GetVarValueFromJS(script, "wanted")

	t.Log(value)
}

func ReplaceJS(js string) string {
	// 替换document[]
	jsContentAfterReplaceDocument := ReplaceDocument(js)

	// 替换setTimeout
	justDoIt := "function just_do_it(p1, p2) { p1(); }"
	finalJS := strings.Replace(jsContentAfterReplaceDocument, "setTimeout", "just_do_it", 1)

	script := "var wanted = 5;" + justDoIt + finalJS

	return script
}

func ReplaceDocument(s string) string {
	// 替换document[]
	jsContentAfterReplaceDocument := s

	documentIndex := strings.Index(s, "document")
	if documentIndex != -1 {
		equalAfterDocumentIndex := strings.Index(s[documentIndex:], "=")
		if equalAfterDocumentIndex != -1 {
			replaceOld := s[documentIndex : documentIndex+equalAfterDocumentIndex]
			jsContentAfterReplaceDocument = strings.Replace(s, replaceOld, "wanted", 1)

			return ReplaceDocument(jsContentAfterReplaceDocument)
		}
	}

	return jsContentAfterReplaceDocument
}

func TestChromedp(t *testing.T) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var res string

	resp, err := chromedp.RunResponse(
		ctx,
		chromedp.Emulate(device.IPadPro),
		chromedp.Navigate(`https://beian.miit.gov.cn/`),
		chromedp.OuterHTML("script", &res),
	)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("status", resp.Headers)

	// t.Log("body", res)
	// nodes := GetScriptNodes(res)
	// js := GetScriptFromScriptNode(nodes[0])
	// htmlutil.WriteToFile("test.js", []byte(js))

}
