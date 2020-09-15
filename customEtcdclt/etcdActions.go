package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/coreos/etcd/clientv3"
)

// etcd action type
const (
	EtcdActNone   = 0
	EtcdActGet    = 1
	EtcdActPut    = 2
	EtcdActDelete = 3
)

// EtcdActionInterface etcd action interface
type EtcdActionInterface interface {
	Exec() ([]string, error)
}

// EtcdActionBase base etcd action
type EtcdActionBase struct {
	ActionType int
}

// EtcdActionGet etcd get action info
type EtcdActionGet struct {
	*EtcdActionBase
	Key      string
	RangeEnd string
}

// EtcdActionPut etcd put action info
type EtcdActionPut struct {
	*EtcdActionBase
	Key   string
	Value string
}

// EtcdActionDelete etcd delete action info
type EtcdActionDelete struct {
	*EtcdActionBase
	Key      string
	RangeEnd string
}

// Exec execute etcd get action
func (action EtcdActionGet) Exec() ([]string, error) {
	if action.Key == "" {
		return nil, errors.New("get command needs one argument as key and an optional argument as range_end")
	}

	client := ConnectEtcd(config)
	if client == nil {
		return nil, errors.New("can not connect to etcd")
	}

	kv := clientv3.NewKV(client)
	var (
		getResp *clientv3.GetResponse
		err     error
	)
	if getResp, err = kv.Get(context.TODO(), action.Key); err != nil {
		return nil, err
	}

	result := make([]string, getResp.Count*2)
	for i, elem := range getResp.Kvs {
		result[i*2] = string(elem.Key)
		result[i*2+1] = string(elem.Value)
	}
	return result, nil
}

// Exec execute etcd put action
func (action EtcdActionPut) Exec() ([]string, error) {
	if action.Key == "" || action.Value == "" {
		return nil, errors.New("put command needs 2 arguments")
	}

	client := ConnectEtcd(config)
	if client == nil {
		return nil, errors.New("can not connect to etcd")
	}

	kv := clientv3.NewKV(client)
	kv.Put(context.TODO(), action.Key, action.Value)
	return []string{"OK"}, nil
}

// Exec execute etcd delete action
func (action EtcdActionDelete) Exec() ([]string, error) {
	if action.Key == "" {
		return nil, errors.New("del command needs one argument as key and an optional argument as range_end")
	}

	client := ConnectEtcd(config)
	if client == nil {
		return nil, errors.New("can not connect to etcd")
	}

	kv := clientv3.NewKV(client)
	var (
		getResp *clientv3.DeleteResponse
		err     error
	)
	if getResp, err = kv.Delete(context.TODO(), action.Key); err != nil {
		return nil, err
	}
	return []string{string(getResp.Deleted)}, nil
}

// ConnectEtcd return etcd client
func ConnectEtcd(config clientv3.Config) *clientv3.Client {
	var (
		client *clientv3.Client = nil
		err    error
	)
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
	}
	return client
}
