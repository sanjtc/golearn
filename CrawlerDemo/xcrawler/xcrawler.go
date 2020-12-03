package xcrawler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pantskun/commonutils/container"
	"golang.org/x/net/html"
)

type Crawler interface {
	Visit(url string)

	// AddHTMLNodeHandler
	// handle html node with filters, filters will be executed in order of input.
	AddHTMLNodeHandler(handler HTMLNodeHandler, filters ...HTMLNodeFilter)
}

type crawler struct {
	rootNode         *html.Node
	htmlNodeHandlers []HTMLNodeHandler
}

var _ Crawler = (*crawler)(nil)

func NewCrawler() Crawler {
	return &crawler{}
}

func (c *crawler) Visit(url string) {
	if err := c.getRootNode(url); err != nil {
		log.Println(err)
		return
	}

	c.traversingAllNode()
}

func (c *crawler) AddHTMLNodeHandler(handler HTMLNodeHandler, filters ...HTMLNodeFilter) {
	h := func(node *html.Node) {
		if !FilterHTMLNode(node, filters...) {
			return
		}

		handler(node)
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
		log.Println("root node is nil")
		return
	}

	queue := container.Queue{}
	queue.Push(c.rootNode)

	for !queue.IsEmpty() {
		curNode := queue.Pop().(*html.Node)

		for _, handler := range c.htmlNodeHandlers {
			handler(curNode)
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
