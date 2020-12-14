package xcrawler

import (
	"log"
	"net/http"
	"net/url"

	"github.com/pantskun/commonutils/container"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
	"golang.org/x/net/html"
)

type Crawler interface {
	Visit(url string)

	// AddHTMLHandler
	// handle html node with filters, filters will be executed in order of input.
	AddHTMLHandler(handler HTMLHandler, filters ...HTMLFilter)

	AddRequestHandler(handler RequestHandler, filters ...RequestFilter)
	AddResponseHandler(handler ResponseHandler, filters ...ResponseFilter)
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

func (c *crawler) Visit(u string) {
	ut, err := url.Parse(u)
	if err != nil {
		xlogutil.Error(err)
		return
	}

	c.visit(ut, 0)
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
	h := func(req *http.Request) {
		if !FilterRequest(req, filters...) {
			return
		}

		handler(req)
	}

	c.requestHandlers = append(c.requestHandlers, h)
}

func (c *crawler) AddResponseHandler(handler ResponseHandler, filters ...ResponseFilter) {
	h := func(resp *http.Response) {
		if !FilterResponse(resp, filters...) {
			return
		}

		handler(resp)
	}

	c.responseHandlers = append(c.responseHandlers, h)
}

func (c *crawler) visit(u *url.URL, depth int) {
	log.Println("visit:", u.String(), ", depth:", depth)

	if c.maxDepth > 0 && depth >= c.maxDepth {
		return
	}

	rootNode := c.getRootNode(u)

	if rootNode == nil {
		xlogutil.Warning("root node is nil, skip")
		return
	}

	rootElement := NewHTMLElement(rootNode, u, depth, c)

	c.traversingAllElement(rootElement)
}

func (c *crawler) getRootNode(u *url.URL) *html.Node {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req := &http.Request{
		Method: "GET",
		URL:    u,
		Header: http.Header{
			"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 Edg/87.0.664.60"},
		},
	}

	for _, handler := range c.requestHandlers {
		handler(req)
	}

	// request经处理后为nil，则不发起请求
	if req.URL == nil {
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		xlogutil.Error(err)
		return nil
	}

	defer resp.Body.Close()

	for _, handler := range c.responseHandlers {
		handler(resp)
	}

	rootNode, err := html.Parse(resp.Body)

	if err != nil {
		xlogutil.Error(err)
		return nil
	}

	return rootNode
}

func (c *crawler) traversingAllElement(rootElement HTMLElement) {
	if rootElement == nil {
		log.Println("warning:", "root node is nil")
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
