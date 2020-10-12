package etcd

import "testing"

func TestEtcdInteraction(t *testing.T) {
	interactor := NewEtcdcltInteractor()
	msg, err := interactor.Get("key")

	if err != nil {
		t.Log(err)
	} else {
		t.Log(msg)
	}
}
