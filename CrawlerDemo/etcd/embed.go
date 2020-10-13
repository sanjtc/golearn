package etcd

import (
	"io/ioutil"
	"os"

	"github.com/coreos/etcd/embed"
)

type embedetcd struct {
	tempdir string
	etcd    *embed.Etcd
}

func newEmbedetcd() *embedetcd {
	tdir, err := ioutil.TempDir(os.TempDir(), "embedetcd")
	if err != nil {
		return nil
	}

	cfg := embed.NewConfig()
	cfg.Dir = tdir

	e, err := embed.StartEtcd(cfg)
	if err != nil {
		return nil
	}

	return &embedetcd{tempdir: tdir, etcd: e}
}

func (e *embedetcd) close() {
	os.RemoveAll(e.tempdir)
	e.etcd.Close()
}
