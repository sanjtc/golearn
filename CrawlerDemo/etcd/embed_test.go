package etcd

import "testing"

func TestEmbedEtcd(t *testing.T) {
	etcd, err := newEmbedetcd()
	if err != nil {
		t.Log(err)
		return
	}

	etcd.close()
}
