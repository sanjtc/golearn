package cmdclient

import (
	"fmt"

	"github.com/pantskun/golearn/customEtcdclt/etcdinteraction"
)

const (
	argc2 = 2
	argc3 = 3
)

// ParseCommand parse command and execute.
func ParseCommand(argv []string) error {
	action := parseCmdAction(argv)
	if action == nil {
		return etcdinteraction.EtcdError{Msg: "action is nil"}
	}

	config := etcdinteraction.ParseEtcdClientConfig("../etcdClientConfig.json")
	if msgs, err := action.Exec(etcdinteraction.GetEtcdClient(config)); err != nil {
		fmt.Println(err)
	} else {
		for _, msg := range msgs {
			fmt.Println(msg)
		}
	}

	return nil
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
	case argc < argc2:
		return etcdinteraction.EtcdActionGet{
			ActionType: etcdinteraction.EtcdActGet,
			Key:        "",
			RangeEnd:   "",
		}
	case argc == argc2:
		return etcdinteraction.EtcdActionGet{
			ActionType: etcdinteraction.EtcdActGet,
			Key:        argv[1],
			RangeEnd:   "",
		}
	case argc > argc2:
		return etcdinteraction.EtcdActionGet{
			ActionType: etcdinteraction.EtcdActGet,
			Key:        argv[1],
			RangeEnd:   argv[2],
		}
	}

	return etcdinteraction.EtcdActionGet{
		ActionType: etcdinteraction.EtcdActGet,
		Key:        "",
		RangeEnd:   "",
	}
}

func parseCmdPutAction(argv []string) etcdinteraction.EtcdActionInterface {
	if argc := len(argv); argc < argc3 {
		return etcdinteraction.EtcdActionPut{
			ActionType: etcdinteraction.EtcdActPut,
			Key:        "",
			Value:      "",
		}
	}

	return etcdinteraction.EtcdActionPut{
		ActionType: etcdinteraction.EtcdActPut,
		Key:        argv[1],
		Value:      argv[2],
	}
}

func parseCmdDeleteAction(argv []string) etcdinteraction.EtcdActionInterface {
	argc := len(argv)

	switch {
	case argc < argc2:
		return etcdinteraction.EtcdActionDelete{
			ActionType: etcdinteraction.EtcdActDelete,
			Key:        "",
			RangeEnd:   "",
		}
	case argc == argc2:
		return etcdinteraction.EtcdActionDelete{
			ActionType: etcdinteraction.EtcdActDelete,
			Key:        argv[1],
			RangeEnd:   "",
		}
	case argc > argc2:
		return etcdinteraction.EtcdActionDelete{
			ActionType: etcdinteraction.EtcdActDelete,
			Key:        argv[1],
			RangeEnd:   argv[2],
		}
	}

	return etcdinteraction.EtcdActionDelete{
		ActionType: etcdinteraction.EtcdActDelete,
		Key:        "",
		RangeEnd:   "",
	}
}
