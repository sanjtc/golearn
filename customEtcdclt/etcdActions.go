package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// etcd action type.
const (
	EtcdActGet    = 1
	EtcdActPut    = 2
	EtcdActDelete = 3
)

const timeoutSecond = 5.0

// EtcdActionInterface etcd action interface.
type EtcdActionInterface interface {
	Equal(EtcdActionInterface) bool
	Exec(client *clientv3.Client) ([]string, error)
}

// EtcdActionGet etcd get action info.
type EtcdActionGet struct {
	ActionType int
	Key        string
	RangeEnd   string
}

// EtcdActionPut etcd put action info.
type EtcdActionPut struct {
	ActionType int
	Key        string
	Value      string
}

// EtcdActionDelete etcd delete action info.
type EtcdActionDelete struct {
	ActionType int
	Key        string
	RangeEnd   string
}

// EtcdError etcd error.
type EtcdError struct {
	msg string
}

func (e EtcdError) Error() string {
	return e.msg
}

// Equal get action.
func (action EtcdActionGet) Equal(b EtcdActionInterface) bool {
	v, ok := b.(EtcdActionGet)
	if !ok || action.Key != v.Key || action.RangeEnd != v.RangeEnd {
		return false
	}

	return true
}

// Equal delete action.
func (action EtcdActionDelete) Equal(b EtcdActionInterface) bool {
	v, ok := b.(EtcdActionDelete)
	if !ok || action.Key != v.Key || action.RangeEnd != v.RangeEnd {
		return false
	}

	return true
}

// Equal put action.
func (action EtcdActionPut) Equal(b EtcdActionInterface) bool {
	v, ok := b.(EtcdActionPut)
	if !ok || action.Key != v.Key || action.Value != v.Value {
		return false
	}

	return true
}

// Exec execute etcd get action.
func (action EtcdActionGet) Exec(client *clientv3.Client) ([]string, error) {
	if action.Key == "" {
		return nil, EtcdError{"get command needs one argument as key and an optional argument as range_end"}
	}

	if client == nil {
		return nil, EtcdError{"can not connect to etcd"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	var (
		getResp *clientv3.GetResponse
		err     error
	)

	if getResp, err = client.Get(ctx, action.Key); err != nil {
		return nil, err
	}

	const TWO = 2

	result := make([]string, getResp.Count*TWO)
	for i, elem := range getResp.Kvs {
		result[i*2] = string(elem.Key)
		result[i*2+1] = string(elem.Value)
	}

	return result, nil
}

// Exec execute etcd put action.
func (action EtcdActionPut) Exec(client *clientv3.Client) ([]string, error) {
	if action.Key == "" || action.Value == "" {
		return nil, EtcdError{"put command needs 2 arguments"}
	}

	if client == nil {
		return nil, EtcdError{"can not connect to etcd"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	if _, err := client.Put(ctx, action.Key, action.Value); err != nil {
		return nil, err
	}

	return []string{"OK"}, nil
}

// Exec execute etcd delete action.
func (action EtcdActionDelete) Exec(client *clientv3.Client) ([]string, error) {
	if action.Key == "" {
		return nil, EtcdError{"del command needs one argument as key and an optional argument as range_end"}
	}

	if client == nil {
		return nil, EtcdError{"can not connect to etcd"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	var (
		getResp *clientv3.DeleteResponse
		err     error
	)

	if getResp, err = client.Delete(ctx, action.Key); err != nil {
		return nil, err
	}

	return []string{strconv.FormatInt(getResp.Deleted, 10)}, nil
}

// GetEtcdClient return etcd client.
func GetEtcdClient() *clientv3.Client {
	config := ParseEtcdClientConfig("etcdClientConfig.json")

	var (
		client *clientv3.Client
		err    error
	)

	if client, err = clientv3.New(config); err != nil {
		log.Println(err)
	}

	return client
}
