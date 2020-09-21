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

		for {
			select {
			case <-ctx.Done():
				return
			case s := <-ss:
				fmt.Println("got signal:", s)
				cancel()

				return
			}
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

	http.HandleFunc("/get", GetActionHandler)
	http.HandleFunc("/put", PutActionHandler)
	http.HandleFunc("/del", DeleteActionHandler)

	fmt.Println("start listen to ", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

// GetActionHandler handle get action.
func GetActionHandler(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query()
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	action := EtcdActionGet{EtcdActGet, key, rangeEnd}
	ExecEtcdAction(action, w)
}

// PutActionHandler handle put action.
func PutActionHandler(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query()
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	value := query.Get("value")
	action := EtcdActionPut{EtcdActPut, key, value}
	ExecEtcdAction(action, w)
}

// DeleteActionHandler handle delete action.
func DeleteActionHandler(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query()
	body, _ := ioutil.ReadAll(r.Body)
	query, _ := url.ParseQuery(string(body))
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	action := EtcdActionDelete{EtcdActDelete, key, rangeEnd}
	ExecEtcdAction(action, w)
}

// ExecEtcdAction exec etcd action and write response.
func ExecEtcdAction(action EtcdActionInterface, w http.ResponseWriter) {
	msgs, err := action.Exec()
	if err != nil {
		_, _ = w.Write([]byte(err.Error() + "\n"))
		return
	}

	for _, msg := range msgs {
		_, _ = w.Write([]byte(msg + "\n"))
	}
}
