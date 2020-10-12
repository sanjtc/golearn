package etcd

import (
	"fmt"
	"os"

	"github.com/coreos/etcd/clientv3"
)

type Interactor interface {
	Get(key string) (string, error)
	Put(key string, value string) error
	Del(key string) error
}

type interactor struct {
	client *clientv3.Client
}

type InteractorError struct {
	msg string
}

func (err *InteractorError) Error() string {
	return fmt.Sprintln(err.msg)
}

func NewInteractor() Interactor {
	os.Getwd()
}
