package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// HTTPClient http client.
func HTTPClient(addr string) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)

	go func() {
		defer wg.Done()
		startHTTPListen(addr, ctx)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		ss := make(chan os.Signal, 1)
		signal.Notify(ss, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return
		case s := <-ss:
			fmt.Println("got signal:", s)
			cancel()

			return
		}
	}()
}

func startHTTPListen(addr string, ctx context.Context) {
	server := &http.Server{Addr: addr, Handler: nil}
	// close the server when ctx done
	go func() {
		<-ctx.Done()
		server.Close()
	}()

	http.HandleFunc("/get", getRequestHandler)
	http.HandleFunc("/put", putRequestHandler)
	http.HandleFunc("/del", deleteRequestHandler)

	fmt.Println("start listen to ", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	action := parseGetRequest(r)
	msg, err := action.Exec(GetEtcdClient())
	writeResponse(msg, err, w)
}

func putRequestHandler(w http.ResponseWriter, r *http.Request) {
	action := parsePutRequest(r)
	msg, err := action.Exec(GetEtcdClient())
	writeResponse(msg, err, w)
}

func deleteRequestHandler(w http.ResponseWriter, r *http.Request) {
	action := parseDeleteRequest(r)
	msg, err := action.Exec(GetEtcdClient())
	writeResponse(msg, err, w)
}

func parseGetRequest(r *http.Request) EtcdActionInterface {
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	return &EtcdActionGet{EtcdActGet, key, rangeEnd}
}

func parsePutRequest(r *http.Request) EtcdActionInterface {
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	value := query.Get("value")

	return &EtcdActionPut{EtcdActPut, key, value}
}

func parseDeleteRequest(r *http.Request) EtcdActionInterface {
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	return &EtcdActionDelete{EtcdActDelete, key, rangeEnd}
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
