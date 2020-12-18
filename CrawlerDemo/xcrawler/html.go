package xcrawler

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type HTMLElement interface {
	Equal(HTMLElement) bool

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

	GetRequest() Request
}

type htmlElement struct {
	node *html.Node
	req  Request
}

func NewHTMLElement(node *html.Node, req Request) HTMLElement {
	return &htmlElement{node: node, req: req}
}

func (e *htmlElement) Equal(other HTMLElement) bool {
	o := other.(*htmlElement)

	return e.node == o.node && e.req == o.req
}

func (e *htmlElement) GetParent() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.Parent == nil {
		return nil
	}

	return &htmlElement{node: e.node.Parent, req: e.req}
}

func (e *htmlElement) GetFirstChild() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.FirstChild == nil {
		return nil
	}

	return &htmlElement{node: e.node.FirstChild, req: e.req}
}

func (e *htmlElement) GetLastChild() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.LastChild == nil {
		return nil
	}

	return &htmlElement{node: e.node.LastChild, req: e.req}
}

func (e *htmlElement) GetPrevSibling() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.PrevSibling == nil {
		return nil
	}

	return &htmlElement{node: e.node.PrevSibling, req: e.req}
}

func (e *htmlElement) GetNextSibling() HTMLElement {
	if e.node == nil {
		return nil
	}

	if e.node.NextSibling == nil {
		return nil
	}

	return &htmlElement{node: e.node.NextSibling, req: e.req}
}

func (e *htmlElement) GetType() html.NodeType {
	if e.node == nil {
		return html.ErrorNode
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

func (e *htmlElement) GetRequest() Request {
	return e.req
}
