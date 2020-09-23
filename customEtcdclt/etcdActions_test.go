package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/coreos/etcd/embed"
	"github.com/coreos/etcd/etcdserver/api/v3client"
)

func TestActionExec(t *testing.T) {
	tdir, err := ioutil.TempDir(os.TempDir(), "embed-test")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tdir)

	// create an embedded EtcdServer from the default configuration
	cfg := embed.NewConfig()
	cfg.Dir = tdir

	e, err := embed.StartEtcd(cfg)
	if err != nil {
		// handle error!
		t.Log(err)
	}
	defer e.Close()

	// wrap the EtcdServer with v3client
	client := v3client.New(e.Server)
	defer client.Close()

	invalidPutActions := []EtcdActionPut{
		{EtcdActPut, "", ""},
		{EtcdActPut, "key1", ""},
		{EtcdActPut, "", "val1"},
	}
	for _, action := range invalidPutActions {
		_, err := action.Exec(client)
		if err == nil {
			t.Errorf("invalid put actions get no error")
		}
	}

	validPutActions := []EtcdActionPut{
		{EtcdActPut, "key1", "val1"},
	}
	for _, action := range validPutActions {
		_, err := action.Exec(client)
		if err != nil {
			t.Log(err)
			t.Errorf("valid put actions get error")
		}
	}

	invalidGetActions := []EtcdActionGet{
		{EtcdActGet, "", ""},
	}
	for _, action := range invalidGetActions {
		_, err := action.Exec(client)
		if err == nil {
			t.Errorf("invalid get actions get no error")
		}
	}

	validGetActions := []EtcdActionGet{
		{EtcdActGet, "key1", ""},
		{EtcdActGet, "key1", "endRange"},
	}
	for _, action := range validGetActions {
		_, err := action.Exec(client)
		if err != nil {
			t.Log(err)
			t.Errorf("valid get actions get error")
		}
	}

	invalidDeleteActions := []EtcdActionDelete{
		{EtcdActDelete, "", ""},
	}
	for _, action := range invalidDeleteActions {
		_, err := action.Exec(client)
		if err == nil {
			t.Errorf("invalid get actions get no error")
		}
	}

	validDeleteActions := []EtcdActionDelete{
		{EtcdActDelete, "key1", ""},
		{EtcdActDelete, "key1", "endRange"},
	}
	for _, action := range validDeleteActions {
		_, err := action.Exec(client)
		if err != nil {
			t.Log(err)
			t.Errorf("valid get actions get error")
		}
	}
}
