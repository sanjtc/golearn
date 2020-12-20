package crawlerutil

import (
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestHrefHandler(t *testing.T) {
	type TestCase struct {
		u        string
		expected bool
	}

	// handleHrefWithAbsolutePathAndHTMLFile
	testHTTPCase := []TestCase{
		{"https://test/test.html", true},
	}
	for _, testCase := range testHTTPCase {
		u, _ := url.Parse(testCase.u)
		got := handleURLWithHTTP(u)
		assert.Equal(t, testCase.expected, got)
	}

	// handleHrefWithJS
	testJSCase := []TestCase{
		{"javascript:void(0)", false},
	}
	for _, testCase := range testJSCase {
		u, _ := url.Parse(testCase.u)
		got := handleURLWithJS(u)
		assert.Equal(t, testCase.expected, got)
	}
}

func TestNewHandlerWithFilters(t *testing.T) {
	type TestCase struct {
		u        string
		handler  URLHandler
		filters  []URLFilter
		expected bool
	}

	testCases := []TestCase{
		{"https://test/test.html", handleURLWithHTTP, []URLFilter{filterURLWithHTTP}, true},
		{"https://test/test.html", handleURLWithHTTP, []URLFilter{filterURLWithJS}, false},
	}

	for _, testCase := range testCases {
		u, _ := url.Parse(testCase.u)
		handler := NewURLHandlerWithFilters(testCase.handler, testCase.filters...)
		got := handler(u)
		assert.Equal(t, testCase.expected, got)
	}
}
