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

	absolutePathTestCases := []TestCase{
		{"https://test/test", true},
		{"http://test/test.html", true},
		{"/test/test/test.html", false},
		{"testtest", false},
	}
	for _, testCase := range absolutePathTestCases {
		u, _ := url.Parse(testCase.href)
		got := FilterHref(u, filterHrefWithAbsolutePath)
		assert.Equal(t, got, testCase.expected)
	}

	relativePathTestCases := []TestCase{
		{"https://test/test", false},
		{"http://test/test.html", false},
		{"/test/test/test.html", true},
		{"testtest", true},
	}
	for _, testCase := range relativePathTestCases {
		u, _ := url.Parse(testCase.href)
		got := FilterHref(u, filterHrefWithRelativePath)
		assert.Equal(t, got, testCase.expected)
	}

	htmlFileTestCases := []TestCase{
		{"https://test/test", false},
		{"http://test/test.html", true},
		{"/test/test/test.html", true},
		{"testtest", false},
	}
	for _, testCase := range htmlFileTestCases {
		u, _ := url.Parse(testCase.href)
		got := FilterHref(u, filterHrefWithFile)
		assert.Equal(t, got, testCase.expected)
	}

	NotFileTestCases := []TestCase{
		{"https://test/test", true},
		{"http://test/test.html", false},
		{"/test/test/test.html", false},
		{"testtest", true},
	}
	for _, testCase := range NotFileTestCases {
		u, _ := url.Parse(testCase.href)
		got := FilterHref(u, filterHrefWithoutFile)
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
		got := FilterHref(u, filterHrefWithJS)
		assert.Equal(t, got, testCase.expected)
	}

	WithQueryTestCase := []TestCase{
		{"https://test/test?a=1", true},
		{"test/test.html?a=1", true},
		{"/test/test/test.html", false},
		{"testtest", false},
	}
	for _, testCase := range WithQueryTestCase {
		u, _ := url.Parse(testCase.href)
		got := FilterHref(u, filterHrefWithQuery)
		assert.Equal(t, got, testCase.expected)
	}

	WithoutQueryTestCase := []TestCase{
		{"https://test/test?a=1", false},
		{"test/test.html?a=1", false},
		{"/test/test/test.html", true},
		{"testtest", true},
	}
	for _, testCase := range WithoutQueryTestCase {
		u, _ := url.Parse(testCase.href)
		got := FilterHref(u, filterHrefWithoutQuery)
		assert.Equal(t, got, testCase.expected)
	}
}
