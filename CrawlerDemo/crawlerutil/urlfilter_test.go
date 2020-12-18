package crawlerutil

import (
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestFilter(t *testing.T) {
	type TestCase struct {
		href     string
		expected bool
	}

	httpTestCases := []TestCase{
		{"https://test/test", true},
		{"http://test/test.html", true},
		{"/test/test/test.html", false},
		{"testtest", false},
	}
	for _, testCase := range httpTestCases {
		u, _ := url.Parse(testCase.href)
		got := FilterURL(u, filterURLWithHTTP)
		assert.Equal(t, got, testCase.expected)
	}

	JSTestCases := []TestCase{
		{"https://test/test", false},
		{"http://test/test.html", false},
		{"/test/test/test.html", false},
		{"testtest", false},
		{"javascript:test", true},
	}
	for _, testCase := range JSTestCases {
		u, _ := url.Parse(testCase.href)
		got := FilterURL(u, filterURLWithJS)
		assert.Equal(t, got, testCase.expected)
	}
}
