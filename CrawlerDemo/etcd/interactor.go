package etcd

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const timeoutSecond = 5.0

type Interactor interface {
	Get(key string) (string, error)
	Put(key string, value string) error
	Del(key string) error
}

type interactor struct {
	client *clientv3.Client
}

var _ Interactor = (*interactor)(nil)

type InteractorError struct {
	msg string
}

func (i *interactor) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	rsp, err := i.client.Get(ctx, key)
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

	_, err := i.client.Put(ctx, key, value)
	if err != nil {
		return err
	}

	return nil
}

func (i *interactor) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel()

	_, err := i.client.Delete(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func (err *InteractorError) Error() string {
	return fmt.Sprintln(err.msg)
}

func NewInteractor() Interactor {
	configFile := GetModulePath() + "configs/etcdConfig.json"
	config := GetClientConfig(configFile)
	client, _ := clientv3.New(config)

	return &interactor{client}
}

func NewInteractorWithClient(client *clientv3.Client) Interactor {
	return &interactor{client}
}

func GetModulePath() string {
	filePath, _ := os.Getwd()
	return path.Dir(filePath)
}
