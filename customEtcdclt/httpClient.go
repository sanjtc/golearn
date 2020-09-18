package main

import (
	"fmt"
	"log"
	"net/http"
)

// StartHTTPListen start http listen.
func StartHTTPListen(addr string) {
	http.HandleFunc("/get", GetActionHandler)
	http.HandleFunc("/put", PutActionHandler)
	http.HandleFunc("/del", DeleteActionHandler)

	fmt.Println("start listen to ", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(err)
	}
}

// GetActionHandler handle get action.
func GetActionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	action := EtcdActionGet{EtcdActGet, key, rangeEnd}
	ExecEtcdAction(action, w)
}

// PutActionHandler handle put action.
func PutActionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	value := query.Get("value")

	action := EtcdActionPut{EtcdActPut, key, value}
	ExecEtcdAction(action, w)
}

// DeleteActionHandler handle delete action.
func DeleteActionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
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
