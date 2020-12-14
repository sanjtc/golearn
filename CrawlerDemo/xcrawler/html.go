package xcrawler

import (
	"net/url"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type HTMLElement interface {
	Equal(HTMLElement) bool
	Visit(u *url.URL)

	GetParent() HTMLElement
	GetFirstChild() HTMLElement
	GetLastChild() HTMLElement
	GetPrevSibling() HTMLElement
	GetNextSibling() HTMLElement
	GetType() html.NodeType
	GetDataAtom() atom.Atom
	GetData() string
	GetNamespace() string
	GetAttr(string) string

	GetOwner() *url.URL
}

type htmlElement struct {
	node  *html.Node
	owner *url.URL

	depth int
	c     *crawler
}

func NewHTMLElement(node *html.Node, owner *url.URL, depth int, c *crawler) HTMLElement {
	return &htmlElement{node: node, owner: owner, depth: depth, c: c}
}

func (e *htmlElement) Equal(other HTMLElement) bool {
	o := other.(*htmlElement)

	return e.node == o.node && e.owner == o.owner
}

func (e *htmlElement) Visit(u *url.URL) {
	e.c.visit(u, e.depth+1)
}

func (e *htmlElement) GetParent() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.Parent == nil {
		return nil
	}

	return &htmlElement{node: e.node.Parent, owner: e.owner, depth: e.depth, c: e.c}
}

func (e *htmlElement) GetFirstChild() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.FirstChild == nil {
		return nil
	}

	return &htmlElement{node: e.node.FirstChild, owner: e.owner, depth: e.depth, c: e.c}
}

func (e *htmlElement) GetLastChild() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.LastChild == nil {
		return nil
	}

	return &htmlElement{node: e.node.LastChild, owner: e.owner, depth: e.depth, c: e.c}
}

func (e *htmlElement) GetPrevSibling() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.PrevSibling == nil {
		return nil
	}

	return &htmlElement{node: e.node.PrevSibling, owner: e.owner, depth: e.depth, c: e.c}
}

func (e *htmlElement) GetNextSibling() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.NextSibling == nil {
		return nil
	}

	return &htmlElement{node: e.node.NextSibling, owner: e.owner, depth: e.depth, c: e.c}
}

func (e *htmlElement) GetType() html.NodeType {
	if e.node == nil {
		return 0
	}

	return e.node.Type
}

func (e *htmlElement) GetDataAtom() atom.Atom {
	if e.node == nil {
		return 0
	}

	return e.node.DataAtom
}

func (e *htmlElement) GetData() string {
	if e.node == nil {
		return ""
	}

	return e.node.Data
}

func (e *htmlElement) GetNamespace() string {
	if e.node == nil {
		return ""
	}

	return e.node.Namespace
}

func (e *htmlElement) GetAttr(key string) string {
	if e.node == nil {
		return ""
	}

	for _, attr := range e.node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}

func (e *htmlElement) GetOwner() *url.URL {
	return e.owner
}
