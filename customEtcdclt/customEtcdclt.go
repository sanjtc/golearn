package main

import (
	"time"

	"github.com/coreos/etcd/clientv3"
)

var config clientv3.Config = clientv3.Config{
	Endpoints:   []string{"127.0.0.1:2379"},
	DialTimeout: 5 * time.Second,
}

func main() {
	// flag.Parse()
	// argv := flag.Args()
	// ParseCommand(argv)
	StartHTTPListen(":8080")
}
