package main

import (
	"flag"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const httpMode = true

var config clientv3.Config = clientv3.Config{
	Endpoints:   []string{"127.0.0.1:2379"},
	DialTimeout: 5 * time.Second,
}

func main() {
	if httpMode {
		StartHTTPListen(":8080")
	} else {
		flag.Parse()
		argv := flag.Args()
		ParseCommand(argv)
	}
}
