package etcd

import (
	"testing"
)

func TestEtcdMutex(t *testing.T) {
	i, err := NewInteractorWithEmbed()

	if err != nil {
		t.Log(err)
		return
	}

	defer i.Close()

	_, err = i.Lock()
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("locked out")

	go func() {
		_, err := i.Lock()
		if err != nil {
			t.Log(err)
			return
		}

		t.Log("locked in")

		// time.Sleep(5 * time.Second)

		_, err = i.Unlock()
		if err != nil {
			t.Log(err)
			return
		}

		t.Log("unlocked in")
	}()

	_, err = i.Unlock()
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("unlocked out")
}
