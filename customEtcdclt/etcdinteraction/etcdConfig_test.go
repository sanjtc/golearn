package etcdinteraction

import (
	"os"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/tools/go/packages"
)

func TestParseEtcdClientConfig(t *testing.T) {
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

	if config := GetEtcdClientConfig(""); !EqualConfig(config, defaultConfig) {
		t.Error("exepected: ", defaultConfig, "got: ", config)
	}

	if config := GetEtcdClientConfig("../etcdClientConfig.json"); EqualConfig(config, defaultConfig) {
		t.Error("got default config")
	}

	if config := GetEtcdClientConfig("../errorEtcdClientConfig.json"); !EqualConfig(config, defaultConfig) {
		t.Error("exepected: ", defaultConfig, "got: ", config)
	}

	if config := GetEtcdClientConfig("../errorEtcdClientConfig2.json"); !EqualConfig(config, defaultConfig) {
		t.Error("exepected: ", defaultConfig, "got: ", config)
	}
}

func TestTest(t *testing.T) {
	cfg := new(packages.Config)
	cfg.Env = os.Environ()
	cfg.Mode = packages.LoadAllSyntax
	// cfg.BuildFlags = []string{"-tags", "fuzz"}

	pkgs, err := packages.Load(cfg, "./")
	if err != nil {
		t.Fatal(err)
	}

	// t.Log(pkgs)

	for _, pkg := range pkgs {
		t.Log(pkg.Syntax)
		s := pkg.Types.Scope()

		for _, n := range s.Names() {
			t.Log(n)
		}
	}
}
