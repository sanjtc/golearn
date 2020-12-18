package xcrawler

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Request interface {
	IsValid() bool
	GetDepth() int
	GetRawReq() *http.Request
	Visit(URL string)
}

type request struct {
	rawReq *http.Request
	depth  int
	c      Crawler
}

func NewRequest(rawReq *http.Request, depth int, c Crawler) Request {
	return &request{rawReq: rawReq, depth: depth, c: c}
}

func (req *request) IsValid() bool {
	return req.rawReq != nil && req.rawReq.URL != nil
}

func (req *request) GetDepth() int {
	return req.depth
}

func (req *request) GetRawReq() *http.Request {
	return req.rawReq
}

func (req *request) Visit(URL string) {
	u, _ := url.Parse(URL)
	req.c.visit(u, req.depth+1)
}

type Response interface {
	GetStatusCode() int
	GetStatus() string
	GetRequest() Request
	GetBody() []byte

	Abandon()
}

type response struct {
	rawResp *http.Response
	request Request
	body    *[]byte

	abandoned bool
}

func NewResponse(rawResp *http.Response, req Request, body *[]byte) Response {
	return &response{rawResp: rawResp, request: req, body: body}
}

func (resp *response) GetStatusCode() int {
	return resp.rawResp.StatusCode
}

func (resp *response) GetStatus() string {
	return resp.rawResp.Status
}

func (resp *response) GetRequest() Request {
	return resp.request
}

// GetBody
// 调用GetBody后才会将io.ReadCloser的body内容封装进[]byte的body中.
func (resp *response) GetBody() []byte {
	if resp.body == nil {
		if resp.rawResp == nil || resp.rawResp.Body == nil {
			return []byte{}
		}

		body, err := ioutil.ReadAll(resp.rawResp.Body)
		if err != nil {
			return []byte{}
		}

		resp.body = &body
	}

	return *resp.body
}

func (resp *response) Abandon() {
	resp.abandoned = true
}
