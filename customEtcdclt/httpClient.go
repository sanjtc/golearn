package main

import (
	"fmt"
	"net/http"
)

// StartHTTPListen start http listen
func StartHTTPListen(addr string) {
	http.HandleFunc("/get", GetActionHandler)
	http.HandleFunc("/put", PutActionHandler)
	http.HandleFunc("/del", DeleteActionHandler)

	fmt.Println("start listen to ", addr)
	http.ListenAndServe(addr, nil)
}

// GetActionHandler handle get action
func GetActionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	action := EtcdActionGet{&EtcdActionBase{EtcdActGet}, key, rangeEnd}
	ExecEtcdAction(action, w)
}

// PutActionHandler handle put action
func PutActionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	value := query.Get("value")

	action := EtcdActionPut{&EtcdActionBase{EtcdActPut}, key, value}
	ExecEtcdAction(action, w)
}

// DeleteActionHandler handle delete action
func DeleteActionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	rangeEnd := query.Get("rangeEnd")

	action := EtcdActionDelete{&EtcdActionBase{EtcdActGet}, key, rangeEnd}
	ExecEtcdAction(action, w)
}

// ExecEtcdAction exec etcd action and write response
func ExecEtcdAction(action EtcdActionInterface, w http.ResponseWriter) {
	msgs, err := action.Exec()

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		w.Write([]byte("\n"))
		return
	}
	for _, msg := range msgs {
		fmt.Println(msg)
		w.Write([]byte(msg))
		w.Write([]byte("\n"))
	}
}
