package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type httpTestInfo struct {
	query  string
	result string
}

const (
	putErrMsg = "put command needs 2 arguments\n"
	getErrMsg = "get command needs one argument as key and an optional argument as range_end\n"
	delErrMsg = "del command needs one argument as key and an optional argument as range_end\n"
)

func TestGetActionHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost:8080/put", nil)

	putCases := []httpTestInfo{
		{"", putErrMsg},
		{"key=key1", putErrMsg},
		{"key=key1&value=value1", "OK\n"},
	}
	for _, putCase := range putCases {
		r.URL.RawQuery = putCase.query
		PutActionHandler(w, r)

		body, _ := ioutil.ReadAll(w.Body)

		bodystr := string(body)
		if bodystr != putCase.result {
			t.Errorf("expect: %s, get: %s", putCase.result, bodystr)
		}
	}

	r = httptest.NewRequest("GET", "http://localhost:8080/get", nil)

	getCases := []httpTestInfo{
		{"", getErrMsg},
		{"key=key1", "key1\nvalue1\n"},
	}
	for _, getCase := range getCases {
		r.URL.RawQuery = getCase.query
		GetActionHandler(w, r)

		body, _ := ioutil.ReadAll(w.Body)
		if string(body) != getCase.result {
			t.Errorf("expect: %s, get: %s", getCase.result, string(body))
		}
	}

	r = httptest.NewRequest("GET", "http://localhost:8080/del", nil)

	delCases := []httpTestInfo{
		{"", delErrMsg},
		{"key=key1", "1\n"},
	}
	for _, delCase := range delCases {
		r.URL.RawQuery = delCase.query
		DeleteActionHandler(w, r)

		body, _ := ioutil.ReadAll(w.Body)
		if string(body) != delCase.result {
			t.Errorf("expect: %s, get: %s", delCase.result, string(body))
		}
	}
}
