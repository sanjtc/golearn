package etcd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/etcdserver/api/v3client"
	"github.com/pantskun/golearn/CrawlerDemo/pathutils"
)

const timeoutSecond = 5.0

type Interactor interface {
	Get(key string) (string, error)
	Put(key string, value string) error
	Del(key string) error
	Lock() error
	Unlock() error
	Close()
}

type interactor struct {
	e *embedetcd
	c *clientv3.Client
	s *concurrency.Session
	m *concurrency.Mutex
}

var _ Interactor = (*interactor)(nil)

type InteractorError struct {
	msg string
}

func NewInteractorWithEmbed() Interactor {
	e := newEmbedetcd()
	if e == nil {
		return nil
	}

	c := v3client.New(e.etcd.Server)

	// new seesion, new mutex
	s, ce := concurrency.NewSession(c)
	if ce != nil {
		log.Println(ce)
		return nil
	}
	m := concurrency.NewMutex(s, "/my-lock/")

	return &interactor{e, c, s, m}
}

func NewInteractor() Interactor {
	configPath := pathutils.GetModulePath() + "/configs/etcdConfig.json"
	config := GetClientConfig(configPath)

	c, err := clientv3.New(config)
	if err != nil {
		return nil
	}

	// new seesion, new mutex
	s, ce := concurrency.NewSession(c)
	if ce != nil {
		log.Println(ce)
		return nil
	}
	m := concurrency.NewMutex(s, "/my-lock/")

	return &interactor{nil, c, s, m}
}

func (i *interactor) Close() {
	if i.e != nil {
		i.e.close()
	}

	i.c.Close()
	i.s.Close()
}

func (i *interactor) Lock() error {
	ctx, _ := context.WithTimeout(context.Background(), timeoutSecond*time.Second)

	err := i.m.Lock(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (i *interactor) Unlock() error {
	ctx, _ := context.WithTimeout(context.Background(), timeoutSecond*time.Second)

	err := i.m.Unlock(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (i *interactor) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	rsp, ge := i.c.Get(ctx, key)
	if ge != nil {
		return "", ge
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
