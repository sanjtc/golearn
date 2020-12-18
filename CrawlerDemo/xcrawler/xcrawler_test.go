package xcrawler

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type serveHandler struct{}

func (s *serveHandler) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	testBody := []byte(
		`
		<html>
			<head></head>
			<body>
    			<h1>
        			test html
    			</h1>
			</body>
		</html>
		`,
	)
	_, _ = respw.Write(testBody)
}

type serveRedirectHandler struct{}

func (s *serveRedirectHandler) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	respw.Header().Add("Location", "http://localhost:2333")
	respw.WriteHeader(301)
}

func TestHandler(t *testing.T) {
	s := &serveHandler{}

	go func() { _ = http.ListenAndServe("localhost:2333", s) }()

	var (
		htmlHandlerExecute     bool = false
		requestHandlerExecute  bool = false
		responseHandlerExecute bool = false
	)

	requestHandler := func(req Request) {
		requestHandlerExecute = true
	}
	responseHandler := func(resp Response) {
		responseHandlerExecute = true
	}
	htmlHandler := func(e HTMLElement) {
		htmlHandlerExecute = true
	}

	c := NewCrawler(1)
	c.AddHTMLHandler(htmlHandler)
	c.AddRequestHandler(requestHandler)
	c.AddResponseHandler(responseHandler)

	c.Visit("http://localhost:2333")

	assert.True(t, htmlHandlerExecute)
	assert.True(t, requestHandlerExecute)
	assert.True(t, responseHandlerExecute)
}

func TestNilURL(t *testing.T) {
	s := &serveHandler{}

	go func() { _ = http.ListenAndServe("localhost:2333", s) }()

	var (
		htmlHandlerExecute     bool = false
		requestHandlerExecute  bool = false
		responseHandlerExecute bool = false
	)

	requestHandler := func(req Request) {
		req.GetRawReq().URL = nil
		requestHandlerExecute = true
	}
	responseHandler := func(resp Response) {
		responseHandlerExecute = true
	}
	htmlHandler := func(e HTMLElement) {
		htmlHandlerExecute = true
	}

	c := NewCrawler(1)
	c.AddHTMLHandler(htmlHandler)
	c.AddRequestHandler(requestHandler)
	c.AddResponseHandler(responseHandler)

	c.Visit("http://localhost:2333")

	assert.True(t, requestHandlerExecute)
	assert.False(t, responseHandlerExecute)
	assert.False(t, htmlHandlerExecute)
}

func TestRedirect(t *testing.T) {
	s1 := &serveRedirectHandler{}
	s2 := &serveHandler{}

	go func() { _ = http.ListenAndServe("localhost:2233", s1) }()
	go func() { _ = http.ListenAndServe("localhost:2333", s2) }()

	var redirectURL *url.URL

	expectedURL, _ := url.Parse("http://localhost:2333")

	responseHandler := func(resp Response) {
		redirectURL = resp.GetRequest().GetRawReq().URL
	}

	c := NewCrawler(1)
	c.AddResponseHandler(responseHandler)

	c.Visit("http://localhost:2233")

	assert.Equal(t, expectedURL.Host, redirectURL.Host)
	assert.Equal(t, expectedURL.Path, redirectURL.Path)
}

func TestDepth(t *testing.T) {
	s := &serveHandler{}

	go func() { _ = http.ListenAndServe("localhost:2333", s) }()

	var (
		htmlHandlerExecute     bool = false
		requestHandlerExecute  bool = false
		responseHandlerExecute bool = false
	)

	requestHandler := func(req Request) {
		requestHandlerExecute = true
	}
	responseHandler := func(resp Response) {
		responseHandlerExecute = true
	}
	htmlHandler := func(e HTMLElement) {
		htmlHandlerExecute = true
	}

	c := NewCrawler(0)
	c.AddHTMLHandler(htmlHandler)
	c.AddRequestHandler(requestHandler)
	c.AddResponseHandler(responseHandler)

	u, _ := url.Parse("http://localhost:2333")
	c.visit(u, 1)

	assert.False(t, requestHandlerExecute)
	assert.False(t, responseHandlerExecute)
	assert.False(t, htmlHandlerExecute)
}
