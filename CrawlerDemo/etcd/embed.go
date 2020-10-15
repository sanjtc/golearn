package etcd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/coreos/etcd/embed"
)

type embedetcd struct {
	tempdir string
	etcd    *embed.Etcd
}

func newEmbedetcd() (*embedetcd, error) {
	tdir, err := ioutil.TempDir(os.TempDir(), "embedetcd")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	cfg := embed.NewConfig()
	cfg.Dir = tdir

	e, err := embed.StartEtcd(cfg)
	if err != nil {
		return nil, err
	}

	return &embedetcd{tempdir: tdir, etcd: e}, nil
}

func (e *embedetcd) close() {
	os.RemoveAll(e.tempdir)
	e.etcd.Close()
}
