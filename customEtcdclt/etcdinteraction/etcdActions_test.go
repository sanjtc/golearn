package etcdinteraction

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
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

func TestNullClientTest(t *testing.T) {
	getAction := EtcdActionGet{EtcdActGet, "", ""}
	putAction := EtcdActionPut{EtcdActPut, "", ""}
	delAction := EtcdActionDelete{EtcdActDelete, "", ""}

	if _, err := getAction.Exec(nil); err == nil {
		t.Errorf("action execute should get error")
	}

	if _, err := putAction.Exec(nil); err == nil {
		t.Errorf("action execute should get error")
	}

	if _, err := delAction.Exec(nil); err == nil {
		t.Errorf("action execute should get error")
	}
}

func TestGetEtcdClient(t *testing.T) {
	{
		config := clientv3.Config{
			Endpoints:   []string{},
			DialTimeout: 5,
		}
		if cli := GetEtcdClient(config); cli != nil {
			t.Error("should got nil, but got:", cli)
		}
	}

	{
		config := clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		}
		if cli := GetEtcdClient(config); cli == nil {
			t.Error("got nil client")
		}
	}
}

func TestEtcdActionEqual(t *testing.T) {
	type testInfo struct {
		action1 EtcdActionInterface
		action2 EtcdActionInterface
	}

	// equal cases
	equalCases := []testInfo{
		{NewDeleteAction("", ""), NewDeleteAction("", "")},
		{NewPutAction("", ""), NewPutAction("", "")},
		{NewGetAction("", ""), NewGetAction("", "")},
	}

	for _, equalCase := range equalCases {
		if !equalCase.action1.Equal(equalCase.action2) {
			t.Error(equalCase.action1, " not equal ", equalCase.action2)
		}
	}

	// unequal cases
	unequalCases := []testInfo{
		{NewDeleteAction("key", ""), NewDeleteAction("", "")},
		{NewDeleteAction("", "rangeEnd"), NewDeleteAction("", "")},
		{NewDeleteAction("", ""), NewPutAction("", "")},

		{NewGetAction("key", ""), NewGetAction("", "")},
		{NewGetAction("", "rangeEnd"), NewGetAction("", "")},
		{NewGetAction("", ""), NewDeleteAction("", "")},

		{NewPutAction("key", ""), NewPutAction("", "")},
		{NewPutAction("", "value"), NewPutAction("", "")},
		{NewPutAction("", ""), NewGetAction("", "")},
	}
	for _, unequalCase := range unequalCases {
		if unequalCase.action1.Equal(unequalCase.action2) {
			t.Error(unequalCase.action1, " equal ", unequalCase.action2)
		}
	}
}
