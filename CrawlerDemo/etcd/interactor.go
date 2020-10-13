package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3client"
)

const timeoutSecond = 5.0

type Interactor interface {
	Get(key string) (string, error)
	Put(key string, value string) error
	Del(key string) error

	Close()
}

type interactor struct {
	e *embedetcd
	c *clientv3.Client
}

var _ Interactor = (*interactor)(nil)

type InteractorError struct {
	msg string
}

func NewInteractor() Interactor {
	e := newEmbedetcd()
	if e == nil {
		return nil
	}

	c := v3client.New(e.etcd.Server)

	return &interactor{e, c}
}

func (i *interactor) Close() {
	i.e.close()
	i.c.Close()
}

func (i *interactor) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	rsp, err := i.c.Get(ctx, key)
	if err != nil {
		return "", err
	}

	if len(rsp.Kvs) == 0 {
		return "", nil
	}

	return string(rsp.Kvs[0].Value), nil
}

func (i *interactor) Put(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	_, err := i.c.Put(ctx, key, value)
	if err != nil {
		return err
	}

	return nil
}

func (i *interactor) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	_, err := i.c.Delete(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func (err *InteractorError) Error() string {
	return fmt.Sprintln(err.msg)
}
