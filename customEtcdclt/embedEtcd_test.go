package main

import (
	"context"
	"testing"

	"github.com/coreos/etcd/embed"
	"github.com/coreos/etcd/etcdserver/api/v3client"
)

func TestEmbed(t *testing.T) {
	// create an embedded EtcdServer from the default configuration
	cfg := embed.NewConfig()
	cfg.Dir = "default.etcd"

	e, err := embed.StartEtcd(cfg)
	if err != nil {
		// handle error!
		t.Log(err)
	}

	// wrap the EtcdServer with v3client
	cli := v3client.New(e.Server)

	// use like an ordinary clientv3
	resp, err := cli.Put(context.TODO(), "some-key", "it works!")
	if err != nil {
		// handle error!
		t.Log(err)
	}

	t.Log(resp)
}
