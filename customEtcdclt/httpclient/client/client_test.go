package client

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/pantskun/golearn/customEtcdclt/etcdinteraction"
)

func TestParseRequest(t *testing.T) {
	type httpTestInfo struct {
		query          string
		expectedAction etcdinteraction.EtcdActionInterface
	}

	r := httptest.NewRequest("GET", "http://localhost:8080/put", nil)
	putCases := []httpTestInfo{
		{"", etcdinteraction.NewPutAction("", "")},
		{"key=key", etcdinteraction.NewPutAction("key", "")},
		{"key=key&value=value", etcdinteraction.NewPutAction("key", "value")},
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
		{"", etcdinteraction.NewGetAction("", "")},
		{"key=key", etcdinteraction.NewGetAction("key", "")},
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
		{"", etcdinteraction.NewDeleteAction("", "")},
		{"key=key", etcdinteraction.NewDeleteAction("key", "")},
	}

	for _, delCase := range delCases {
		r.Body = ioutil.NopCloser(bytes.NewBufferString(delCase.query))

		action := parseDeleteRequest(r)
		if !action.Equal(delCase.expectedAction) {
			t.Errorf("expect: %s, get: %s", delCase.expectedAction, action)
		}
	}
}

func TestHTTPClient(t *testing.T) {
	wg := sync.WaitGroup{}

	var err error

	wg.Add(1)

	ctrlBreakChan := make(chan os.Signal, 1)

	go func() {
		defer wg.Done()

		err = HTTPClient(":8080", ctrlBreakChan)
	}()

	// {
	// 	sendCtrlBreak := func(t *testing.T, pid int) {
	// 		d, e := syscall.LoadDLL("kernel32.dll")
	// 		if e != nil {
	// 			t.Fatalf("LoadDLL: %v\n", e)
	// 		}

	// 		p, e := d.FindProc("GenerateConsoleCtrlEvent")
	// 		if e != nil {
	// 			t.Fatalf("FindProc: %v\n", e)
	// 		}

	// 		r, _, e := p.Call(syscall.CTRL_BREAK_EVENT, uintptr(pid))
	// 		if r == 0 {
	// 			t.Fatalf("GenerateConsoleCtrlEvent: %v\n", e)
	// 		}
	// 	}

	// 	sendCtrlBreak(t, syscall.Getpid())
	// } // windows

	// {
	// 	time.Sleep(1 * time.Second)
	// 	_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	// } // linux

	time.Sleep(1 * time.Second)

	ctrlBreakChan <- os.Interrupt

	wg.Wait()

	if err.Error() != "http: Server closed" {
		t.Error("not got server close")
	}
}
