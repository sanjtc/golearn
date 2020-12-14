package etcd

import (
	"io/ioutil"
	"os"

	"github.com/coreos/etcd/embed"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

type embedetcd struct {
	tempdir string
	etcd    *embed.Etcd
}

func newEmbedetcd() (*embedetcd, error) {
	tdir, err := ioutil.TempDir(os.TempDir(), "embedetcd")
	if err != nil {
		xlogutil.Error(err)
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
