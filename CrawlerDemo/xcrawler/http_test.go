package xcrawler

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := &crawler{}
	rawReq := &http.Request{URL: &url.URL{}}

	req := NewRequest(rawReq, 0, c)

	assert.True(t, req.IsValid())
	assert.Equal(t, req.GetDepth(), 0)
	assert.Equal(t, req.GetRawReq(), rawReq)
}

func TestResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &request{}
	rawResp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
	}

	resp := NewResponse(rawResp, req, nil)

	assert.Equal(t, 200, resp.GetStatusCode())
	assert.Equal(t, "200 OK", resp.GetStatus())
	assert.Equal(t, req, resp.GetRequest())
	assert.Equal(t, []byte{}, resp.GetBody())
}

func TestVisit(t *testing.T) {
	var (
		htmlHandlerExecute     bool = false
		requestHandlerExecute  bool = false
		responseHandlerExecute bool = false
	)

	htmlHandler := func(e HTMLElement) {
		htmlHandlerExecute = true
	}
	requestHandler := func(req Request) {
		requestHandlerExecute = true
	}
	responseHandler := func(resp Response) {
		responseHandlerExecute = true
	}

	s := &serveHandler{}
	go func() { http.ListenAndServe("localhost:2333", s) }()

	c := NewCrawler(2)
	c.AddHTMLHandler(htmlHandler)
	c.AddRequestHandler(requestHandler)
	c.AddResponseHandler(responseHandler)

	req := &request{depth: 0, c: c}
	req.Visit("http://localhost:2333")

	assert.True(t, htmlHandlerExecute)
	assert.True(t, requestHandlerExecute)
	assert.True(t, responseHandlerExecute)
}
