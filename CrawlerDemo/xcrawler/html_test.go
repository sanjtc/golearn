package xcrawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestHTMLElement(t *testing.T) {
	req := &request{}

	rootNode := &html.Node{
		Type:      html.DocumentNode,
		DataAtom:  atom.A,
		Data:      "testData",
		Namespace: "testNamespace",
		Attr: []html.Attribute{
			{Namespace: "testNamespace", Key: "testKey", Val: "testVal"},
		},
	}
	childNode1 := &html.Node{
		Type:      html.DocumentNode,
		DataAtom:  atom.A,
		Data:      "testData",
		Namespace: "testNamespace",
		Attr: []html.Attribute{
			{Namespace: "testNamespace", Key: "testKey", Val: "testVal"},
		},
	}
	childNode2 := &html.Node{
		Type:      html.DocumentNode,
		DataAtom:  atom.A,
		Data:      "testData",
		Namespace: "testNamespace",
		Attr: []html.Attribute{
			{Namespace: "testNamespace", Key: "testKey", Val: "testVal"},
		},
	}

	rootNode.FirstChild = childNode1
	rootNode.LastChild = childNode2
	childNode1.Parent = rootNode
	childNode1.NextSibling = childNode2
	childNode2.Parent = rootNode
	childNode2.PrevSibling = childNode1

	nilElement := &htmlElement{}
	rootElement := NewHTMLElement(rootNode, req)
	childElement1 := NewHTMLElement(childNode1, req)
	childElement2 := NewHTMLElement(childNode2, req)

	//  Test GetParent
	assert.True(t, rootElement.Equal(childElement1.GetParent()))
	assert.True(t, rootElement.Equal(childElement2.GetParent()))
	assert.Equal(t, rootElement.GetParent(), nil)
	assert.Equal(t, nilElement.GetParent(), nil)

	// Test GetFirstChild
	assert.True(t, childElement1.Equal(rootElement.GetFirstChild()))
	assert.Equal(t, childElement1.GetFirstChild(), nil)
	assert.Equal(t, nilElement.GetFirstChild(), nil)

	// Test GetPrevSibling
	assert.True(t, childElement1.Equal(childElement2.GetPrevSibling()))
	assert.Equal(t, rootElement.GetPrevSibling(), nil)
	assert.Equal(t, nilElement.GetPrevSibling(), nil)

	// Test GetLastChild
	assert.True(t, childElement2.Equal(rootElement.GetLastChild()))
	assert.Equal(t, childElement2.GetLastChild(), nil)
	assert.Equal(t, nilElement.GetLastChild(), nil)

	// Test GetNextSibling
	assert.True(t, childElement2.Equal(childElement1.GetNextSibling()))
	assert.Equal(t, rootElement.GetNextSibling(), nil)
	assert.Equal(t, nilElement.GetNextSibling(), nil)

	// Test GetType
	assert.Equal(t, rootElement.GetType(), html.DocumentNode)
	assert.Equal(t, nilElement.GetType(), html.ErrorNode)

	// Test GetDataAtom
	assert.Equal(t, rootElement.GetDataAtom(), atom.A)
	assert.Equal(t, nilElement.GetDataAtom(), atom.Atom(0))

	// Test GetData
	assert.Equal(t, rootElement.GetData(), "testData")
	assert.Equal(t, nilElement.GetData(), "")

	// Test GetNamespace
	assert.Equal(t, rootElement.GetNamespace(), "testNamespace")
	assert.Equal(t, nilElement.GetNamespace(), "")

	// Test GetAttr
	assert.Equal(t, rootElement.GetAttr("testKey"), "testVal")
	assert.Equal(t, rootElement.GetAttr("none"), "")
	assert.Equal(t, nilElement.GetAttr("none"), "")

	// Test GetRequest
	assert.Equal(t, rootElement.GetRequest(), req)
}
