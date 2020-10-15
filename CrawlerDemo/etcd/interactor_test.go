package etcd

import "testing"

func TestEtcdMutex(t *testing.T) {
	i, err := NewInteractor()

	if err != nil {
		return
	}

	i.Lock()
}
