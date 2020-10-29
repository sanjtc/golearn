package client

import (
	"flag"
	"path"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/customEtcdclt/etcdinteraction"
)

const (
	argc2 = 2
)

func CMDClient() (string, error) {
	flag.Parse()
	argv := flag.Args()
	modulePath := pathutils.GetModulePath("customEtcdclt")
	config := etcdinteraction.GetEtcdClientConfig(path.Join(modulePath, "configs/etcdClientConfig.json"))

	return etcdinteraction.ExecuteAction(parseCmdAction(argv), etcdinteraction.GetEtcdClient(config))
}

func parseCmdAction(argv []string) etcdinteraction.EtcdActionInterface {
	if argc := len(argv); argc == 0 {
		return nil
	}

	command := argv[0]
	switch command {
	case "get":
		return parseCmdGetAction(argv)
	case "put":
		return parseCmdPutAction(argv)
	case "del":
		return parseCmdDeleteAction(argv)
	}

	return nil
}

func parseCmdGetAction(argv []string) etcdinteraction.EtcdActionInterface {
	argc := len(argv)

	switch {
	case argc == argc2:
		return etcdinteraction.NewGetAction(argv[1], "")
	case argc > argc2:
		return etcdinteraction.NewGetAction(argv[1], argv[2])
	default:
		return etcdinteraction.NewGetAction("", "")
	}
}

func parseCmdPutAction(argv []string) etcdinteraction.EtcdActionInterface {
	argc := len(argv)

	switch {
	case argc == argc2:
		return etcdinteraction.NewPutAction(argv[1], "")
	case argc > argc2:
		return etcdinteraction.NewPutAction(argv[1], argv[2])
	default:
		return etcdinteraction.NewPutAction("", "")
	}
}

func parseCmdDeleteAction(argv []string) etcdinteraction.EtcdActionInterface {
	argc := len(argv)

	switch {
	case argc == argc2:
		return etcdinteraction.NewDeleteAction(argv[1], "")
	case argc > argc2:
		return etcdinteraction.NewDeleteAction(argv[1], argv[2])
	default:
		return etcdinteraction.NewDeleteAction("", "")
	}
}
