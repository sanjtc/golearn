package crawler

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestGetElementNodesFromURL(t *testing.T) {
	type TestCase struct {
		url      string
		element  string
		expected string
	}

	testCases := []TestCase{
		{url: "https://www.ssetech.com.cn/", element: "a", expected: "a"},
	}

	for _, testCase := range testCases {
		nodes := GetElementNodesFromURL(testCase.url, testCase.element)
		for _, node := range nodes {
			assert.Equal(t, node.Data, testCase.expected)
		}
	}
}

func TestGetElementAttributeValue(t *testing.T) {
	type TestCase struct {
		element   *html.Node
		attribute string
		expected  string
	}

	element1 := &html.Node{
		Attr: []html.Attribute{
			{Namespace: "", Key: "key", Val: "value"},
		},
	}

	testCases := []TestCase{
		{element: element1, attribute: "key", expected: "value"},
		{element: element1, attribute: "error", expected: ""},
		{element: nil, attribute: "", expected: ""},
	}

	for _, testCase := range testCases {
		value := GetElementAttributeValue(testCase.element, testCase.attribute)
		assert.Equal(t, value, testCase.expected)
	}
}

func TestFilterURL(t *testing.T) {
	urls := []string{
		":",
		"http://test1",
		"https://test2",
	}

	filter1 := func(s string) bool {
		return strings.HasPrefix(s, "http")
	}

	urls = FilterURL(urls, filter1)

	for _, url := range urls {
		assert.Equal(t, strings.HasPrefix(url, "http"), true)
	}
}

func TestDownloadURL(t *testing.T) {
	modulePath := pathutils.GetModulePath("CrawlerDemo")
	downloadPath := path.Join(modulePath, "crawler/temp")

	defer os.RemoveAll(downloadPath)

	url := "https://www.baidu.com/index.html"

	err := DownloadURL(url, downloadPath)
	assert.Nil(t, err)
}
