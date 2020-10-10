package httpclient

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"

	"github.com/pantskun/golearn/customEtcdclt/etcdinteraction"
)

// HTTPClient http client.
func HTTPClient(addr string, ctrlBreakChan chan os.Signal) error {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)

	go func() {
		defer wg.Done()

		listenSystemSignal(ctrlBreakChan, ctx, cancel)
	}()

	return startHTTPListen(addr, ctx)
}

func listenSystemSignal(ctrlBreakChan chan os.Signal, ctx context.Context, cancel context.CancelFunc) {
	signal.Notify(ctrlBreakChan)

	for {
		select {
		case <-ctx.Done():
			return
		case s := <-ctrlBreakChan:
			if s == os.Interrupt {
				fmt.Println("got signal:", s)
				cancel()

				return
			}
		}
	}
}

func startHTTPListen(addr string, ctx context.Context) error {
	server := &http.Server{Addr: addr, Handler: nil}
	// close server when ctx done
	go func() {
		<-ctx.Done()
		server.Close()
	}()

	http.HandleFunc("/get", getRequestHandler)
	http.HandleFunc("/put", putRequestHandler)
	http.HandleFunc("/del", deleteRequestHandler)

	fmt.Println("start listen to ", addr)

	return server.ListenAndServe()
}

func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	action := parseGetRequest(r)
	execActionAndWriteResponse(action, w)
}

func putRequestHandler(w http.ResponseWriter, r *http.Request) {
	action := parsePutRequest(r)
	execActionAndWriteResponse(action, w)
}

func deleteRequestHandler(w http.ResponseWriter, r *http.Request) {
	action := parseDeleteRequest(r)
	execActionAndWriteResponse(action, w)
}

func execActionAndWriteResponse(action etcdinteraction.EtcdActionInterface, w http.ResponseWriter) {
	config := etcdinteraction.GetEtcdClientConfig("../../etcdClientConfig.json")
	msg, err := action.Exec(etcdinteraction.GetEtcdClient(config))
	writeResponse(msg, err, w)
}

func parseGetRequest(r *http.Request) etcdinteraction.EtcdActionInterface {
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	return etcdinteraction.NewGetAction(key, rangeEnd)
}

func parsePutRequest(r *http.Request) etcdinteraction.EtcdActionInterface {
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	value := query.Get("value")

	return etcdinteraction.NewPutAction(key, value)
}

func parseDeleteRequest(r *http.Request) etcdinteraction.EtcdActionInterface {
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	return etcdinteraction.NewDeleteAction(key, rangeEnd)
}

func writeResponse(msgs []string, err error, w http.ResponseWriter) {
	if err != nil {
		_, _ = w.Write([]byte(err.Error() + "\n"))
		return
	}

	for _, msg := range msgs {
		_, _ = w.Write([]byte(msg + "\n"))
	}
}
