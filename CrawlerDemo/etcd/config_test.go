package etcd

import (
	"path"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pantskun/commonutils/pathutils"
)

func TestGetClientConfig(t *testing.T) {
	const timeSecond = 5.0

	defaultConfig := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: timeSecond * time.Second,
	}

	EqualConfig := func(config1, config2 clientv3.Config) bool {
		if config1.DialTimeout != config2.DialTimeout {
			return false
		}

		endpointNum1 := len(config1.Endpoints)
		endpointNum2 := len(config2.Endpoints)

		if endpointNum1 != endpointNum2 {
			return false
		}

		for i := 0; i < endpointNum1; i++ {
			if config1.Endpoints[i] != config2.Endpoints[i] {
				return false
			}
		}

		return true
	}

	modulePath := pathutils.GetModulePath("CrawlerDemo")

	if config := GetClientConfig(""); !EqualConfig(config, defaultConfig) {
		t.Error("exepected: ", defaultConfig, "got: ", config)
	}

	if config := GetClientConfig(path.Join(modulePath, "configs/etcdClientConfig.json")); EqualConfig(config, defaultConfig) {
		t.Error("got default config")
	}

	if config := GetClientConfig(path.Join(modulePath, "configs/errorEtcdClientConfig.json")); !EqualConfig(config, defaultConfig) {
		t.Error("exepected: ", defaultConfig, "got: ", config)
	}

	if config := GetClientConfig(path.Join(modulePath, "configs/errorEtcdClientConfig2.json")); !EqualConfig(config, defaultConfig) {
		t.Error("exepected: ", defaultConfig, "got: ", config)
	}
}
