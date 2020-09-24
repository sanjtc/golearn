package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/pantskun/golearn/customEtcdclt/etcdinteraction"
)

// const (
// 	putErrMsg = "put command needs 2 arguments\n"
// 	getErrMsg = "get command needs one argument as key and an optional argument as range_end\n"
// 	delErrMsg = "del command needs one argument as key and an optional argument as range_end\n"
// )

func TestParseRequest(t *testing.T) {
	type httpTestInfo struct {
		query          string
		expectedAction etcdinteraction.EtcdActionInterface
	}

	r := httptest.NewRequest("GET", "http://localhost:8080/put", nil)
	putCases := []httpTestInfo{
		{"", etcdinteraction.EtcdActionPut{ActionType: etcdinteraction.EtcdActPut, Key: "", Value: ""}},
		{"key=key", etcdinteraction.EtcdActionPut{ActionType: etcdinteraction.EtcdActPut, Key: "key", Value: ""}},
		{"key=key&value=value", etcdinteraction.EtcdActionPut{ActionType: etcdinteraction.EtcdActPut, Key: "key", Value: "value"}},
	}

	for _, putCase := range putCases {
		r.Body = ioutil.NopCloser(bytes.NewBufferString(putCase.query))

		action := parsePutRequest(r)
		if !action.Equal(putCase.expectedAction) {
			t.Errorf("expect: %s, get: %s", putCase.expectedAction, action)
		}
	}

	r = httptest.NewRequest("GET", "http://localhost:8080/get", nil)
	getCases := []httpTestInfo{
		{"", etcdinteraction.EtcdActionGet{ActionType: etcdinteraction.EtcdActGet, Key: "", RangeEnd: ""}},
		{"key=key", etcdinteraction.EtcdActionGet{ActionType: etcdinteraction.EtcdActGet, Key: "key", RangeEnd: ""}},
	}

	for _, getCase := range getCases {
		r.Body = ioutil.NopCloser(bytes.NewBufferString(getCase.query))

		action := parseGetRequest(r)
		if !action.Equal(getCase.expectedAction) {
			t.Errorf("expect: %s, get: %s", getCase.expectedAction, action)
		}
	}

	r = httptest.NewRequest("GET", "http://localhost:8080/del", nil)
	delCases := []httpTestInfo{
		{"", etcdinteraction.EtcdActionDelete{ActionType: etcdinteraction.EtcdActDelete, Key: "", RangeEnd: ""}},
		{"key=key", etcdinteraction.EtcdActionDelete{ActionType: etcdinteraction.EtcdActDelete, Key: "key", RangeEnd: ""}},
	}

	for _, delCase := range delCases {
		r.Body = ioutil.NopCloser(bytes.NewBufferString(delCase.query))

		action := parseDeleteRequest(r)
		if !action.Equal(delCase.expectedAction) {
			t.Errorf("expect: %s, get: %s", delCase.expectedAction, action)
		}
	}
}

func TestWriteResponse(t *testing.T) {
	type writeTestInfo struct {
		msgs     []string
		err      error
		expected string
	}

	w := httptest.NewRecorder()

	writeCases := []writeTestInfo{
		{nil, etcdinteraction.EtcdError{Msg: "test"}, "test\n"},
		{[]string{"test1", "test2"}, nil, "test1\ntest2\n"},
	}

	for _, writeCase := range writeCases {
		writeResponse(writeCase.msgs, writeCase.err, w)

		body, _ := ioutil.ReadAll(w.Body)
		bodystr := string(body)

		if bodystr != writeCase.expected {
			t.Errorf("expect: %s, get: %s", writeCase.expected, bodystr)
		}
	}
}
