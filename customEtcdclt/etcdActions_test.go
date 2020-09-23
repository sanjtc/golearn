package main

import (
	"io/ioutil"
	"os"
	"testing"

	embedEtcd "github.com/coreos/etcd/embed"
)

func TestActionExec(t *testing.T) {
	invalidPutActions := []EtcdActionPut{
		{EtcdActPut, "", ""},
		{EtcdActPut, "key1", ""},
		{EtcdActPut, "", "val1"},
	}
	for _, action := range invalidPutActions {
		_, err := action.Exec()
		if err == nil {
			t.Errorf("invalid put actions get no error")
		}
	}

	validPutActions := []EtcdActionPut{
		{EtcdActPut, "key1", "val1"},
	}
	for _, action := range validPutActions {
		_, err := action.Exec()
		if err != nil {
			t.Errorf("valid put actions get error")
		}
	}

	invalidGetActions := []EtcdActionGet{
		{EtcdActGet, "", ""},
	}
	for _, action := range invalidGetActions {
		_, err := action.Exec()
		if err == nil {
			t.Errorf("invalid get actions get no error")
		}
	}

	validGetActions := []EtcdActionGet{
		{EtcdActGet, "key1", ""},
		{EtcdActGet, "key1", "endRange"},
	}
	for _, action := range validGetActions {
		_, err := action.Exec()
		if err != nil {
			t.Errorf("valid get actions get error")
		}
	}

	invalidDeleteActions := []EtcdActionDelete{
		{EtcdActDelete, "", ""},
	}
	for _, action := range invalidDeleteActions {
		_, err := action.Exec()
		if err == nil {
			t.Errorf("invalid get actions get no error")
		}
	}

	validDeleteActions := []EtcdActionDelete{
		{EtcdActDelete, "key1", ""},
		{EtcdActDelete, "key1", "endRange"},
	}
	for _, action := range validDeleteActions {
		_, err := action.Exec()
		if err != nil {
			t.Errorf("valid get actions get error")
		}
	}
}

func TestEmbedEtcd(t *testing.T) {
	tdir, err := ioutil.TempDir(os.TempDir(), "auth-test")
	if err != nil {
		t.Fatal(err)
	}
	cfg := embedEtcd.NewConfig()
	cfg.Dir = tdir
	embedEtcd.StartEtcd(cfg)
}
