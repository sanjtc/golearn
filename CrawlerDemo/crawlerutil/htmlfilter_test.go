package crawlerutil

import (
	"testing"

	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestFilterANode(t *testing.T) {
	ANode := html.Node{Data: "a"}
	OtherNode1 := html.Node{Data: "h"}
	OtherNode2 := html.Node{Data: ""}

	type TestCase struct {
		node     *html.Node
		expected bool
	}

	testCases := []TestCase{
		{node: &ANode, expected: true},
		{node: &OtherNode1, expected: false},
		{node: &OtherNode2, expected: false},
	}

	for _, testCase := range testCases {
		got := xcrawler.FilterHTML(testCase.node, FilterANode)
		assert.Equal(t, testCase.expected, got)
	}
}
