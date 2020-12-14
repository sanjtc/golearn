package crawlerutil

import (
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

// SynchronizeWithETCD
// if this key is not in etcd, put key to etcd and return true, otherwise return false.
func Synchronize(key string, etcdInteractor etcd.Interactor) bool {
	if etcdInteractor == nil {
		return true
	}

	// lock
	if _, err := etcdInteractor.Lock(); err != nil {
		xlogutil.Error(err)
		return false
	}

	defer func() {
		// unlock
		if _, err := etcdInteractor.Unlock(); err != nil {
			xlogutil.Error(err)
		}
	}()

	// check url
	res, err := etcdInteractor.Get(key)
	if err != nil {
		xlogutil.Error(err)
		return false
	}

	if res == "" {
		err := etcdInteractor.Put(key, "1")
		if err != nil {
			xlogutil.Error(err)
			return false
		}

		return true
	}

	return false
}
