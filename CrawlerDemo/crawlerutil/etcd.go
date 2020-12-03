package crawlerutil

import (
	"log"

	"github.com/pantskun/golearn/CrawlerDemo/etcd"
)

// SynchronizeWithETCD
// if this key is not in etcd, put key to etcd and return true, otherwise return false.
// if this key is not in etcd, put key to etcd and return true, otherwise return false.
func Synchronize(key string, etcdInteractor etcd.Interactor) bool {
	// lock
	if _, err := etcdInteractor.Lock(); err != nil {
		log.Println(err)
		return false
	}

	defer func() {
		// unlock
		if _, err := etcdInteractor.Unlock(); err != nil {
			log.Println(err)
		}
	}()

	// check url
	res, err := etcdInteractor.Get(key)
	if err != nil {
		log.Println(err)
		return false
	}

	if res == "" {
		err := etcdInteractor.Put(key, "1")
		if err != nil {
			log.Println(err)
			return false
		}

		return true
	}

	return false
}