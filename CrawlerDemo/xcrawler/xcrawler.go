package xcrawler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pantskun/commonutils/container"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
	"golang.org/x/net/html"
)

type Crawler interface {
	Visit(url string)
	GetHost() string

	// AddHTMLHandler
	// handle html node with filters, filters will be executed in order of input.
	AddHTMLHandler(handler HTMLHandler, filters ...HTMLFilter)
}

type crawler struct {
	rawURL           string
	host             string
	rootNode         *html.Node
	htmlNodeHandlers []HTMLHandler
}

var _ Crawler = (*crawler)(nil)

func NewCrawler() Crawler {
	return &crawler{}
}

func (c *crawler) Visit(url string) {
	c.rawURL = url
	c.host = htmlutil.GetDomainFromURL(url)

	if err := c.getRootNode(url); err != nil {
		log.Println("error:", err)
		return
	}

	c.traversingAllNode()
}

func (c *crawler) GetHost() string {
	return c.host
}

func (c *crawler) AddHTMLHandler(handler HTMLHandler, filters ...HTMLFilter) {
	h := func(node *html.Node, c Crawler) {
		if !FilterHTML(node, filters...) {
			return
		}

		handler(node, c)
	}

	c.htmlNodeHandlers = append(c.htmlNodeHandlers, h)
}

func (c *crawler) getRootNode(url string) error {
	resp, err := http.Get(fmt.Sprint(url))
	if err != nil {
		return err
	}

	c.rootNode, err = html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return err
	}

	return nil
}

func (c *crawler) traversingAllNode() {
	if c.rootNode == nil {
		log.Println("warning:", "root node is nil")
		return
	}

	queue := container.Queue{}
	queue.Push(c.rootNode)

	for !queue.IsEmpty() {
		curNode := queue.Pop().(*html.Node)

		for _, handler := range c.htmlNodeHandlers {
			handler(curNode, c)
		}

		childNode := curNode.FirstChild
		if childNode == nil {
			continue
		}

		for childNode != curNode.LastChild {
			queue.Push(childNode)
			childNode = childNode.NextSibling
		}

		queue.Push(childNode)
	}
}
