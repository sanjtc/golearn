package xcrawler

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/pantskun/commonutils/container"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
	"golang.org/x/net/html"
)

type Crawler interface {
	Visit(URL string)
	AddHTMLHandler(handler HTMLHandler, filters ...HTMLFilter)
	AddRequestHandler(handler RequestHandler, filters ...RequestFilter)
	AddResponseHandler(handler ResponseHandler, filters ...ResponseFilter)

	visit(URL *url.URL, depth int)
}

type crawler struct {
	maxDepth int

	htmlHandlers     []HTMLHandler
	requestHandlers  []RequestHandler
	responseHandlers []ResponseHandler
}

var _ Crawler = (*crawler)(nil)

func NewCrawler(maxDepth int) Crawler {
	return &crawler{maxDepth: maxDepth}
}

func (c *crawler) Visit(URL string) {
	u, err := url.Parse(URL)
	if err != nil {
		xlogutil.Error(err)
		return
	}

	c.visit(u, 0)
}

func (c *crawler) AddHTMLHandler(handler HTMLHandler, filters ...HTMLFilter) {
	h := func(element HTMLElement) {
		if !FilterHTML(element, filters...) {
			return
		}

		handler(element)
	}

	c.htmlHandlers = append(c.htmlHandlers, h)
}

func (c *crawler) AddRequestHandler(handler RequestHandler, filters ...RequestFilter) {
	h := func(req Request) {
		if !FilterRequest(req, filters...) {
			return
		}

		handler(req)
	}

	c.requestHandlers = append(c.requestHandlers, h)
}

func (c *crawler) AddResponseHandler(handler ResponseHandler, filters ...ResponseFilter) {
	h := func(resp Response) {
		if !FilterResponse(resp, filters...) {
			return
		}

		handler(resp)
	}

	c.responseHandlers = append(c.responseHandlers, h)
}

func (c *crawler) visit(u *url.URL, depth int) {
	xlogutil.Warning("visit:", u.String(), ", depth:", depth)

	if c.maxDepth >= 0 && depth > c.maxDepth {
		xlogutil.Warning("reach max depth, stop")
		return
	}

	rootElement := c.getRootNode(u, depth)

	c.traversingAllElement(rootElement)
}

func (c *crawler) getRootNode(u *url.URL, depth int) HTMLElement {
	var lastRawReq *http.Request

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			lastRawReq = req
			return nil
		},
	}

	rawReq := &http.Request{
		Method: "GET",
		URL:    u,
		Header: http.Header{
			"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 Edg/87.0.664.60"},
		},
	}
	// 封装http.Request
	req := &request{rawReq: rawReq, depth: depth, c: c}

	for _, handler := range c.requestHandlers {
		handler(req)
	}

	// request经处理后为nil，则不发起请求(client.Do中对request的url已经进行了错误检查)
	// if req == nil || !req.IsValid() {
	// 	return nil
	// }

	rawResp, err := client.Do(req.rawReq)
	if err != nil {
		xlogutil.Error(err)
		return nil
	}
	defer rawResp.Body.Close()

	// 发生重定向时，response的request为最后一次请求的request
	if lastRawReq != nil {
		req.rawReq = lastRawReq
	}

	// 封装http.Response，body不需要立即读取，调用GetBody时进行读取
	resp := &response{rawResp: rawResp, request: req, body: nil}

	for _, handler := range c.responseHandlers {
		handler(resp)
	}

	if resp.abandoned {
		return nil
	}

	rootNode, err := html.Parse(bytes.NewReader(resp.GetBody()))
	if err != nil {
		xlogutil.Error(err)
		return nil
	}

	return NewHTMLElement(rootNode, req)
}

func (c *crawler) traversingAllElement(rootElement HTMLElement) {
	if rootElement == nil {
		xlogutil.Warning("root node is nil")
		return
	}

	queue := container.Queue{}
	queue.Push(rootElement)

	for !queue.IsEmpty() {
		curElement := queue.Pop().(HTMLElement)

		for _, handler := range c.htmlHandlers {
			handler(curElement)
		}

		childElement := curElement.GetFirstChild()
		if childElement == nil {
			continue
		}

		for !childElement.Equal(curElement.GetLastChild()) {
			queue.Push(childElement)
			childElement = childElement.GetNextSibling()
		}

		queue.Push(childElement)
	}
}
