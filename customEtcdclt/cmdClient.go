package main

import (
	"fmt"
)

const (
	argc2 = 2
	argc3 = 3
)

// ParseCommand parse command and execute.
func ParseCommand(argv []string) {
	action := parseCmdAction(argv)
	if action == nil {
		return
	}

	if msgs, err := action.Exec(); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, msg := range msgs {
			fmt.Println(msg)
		}
	}
}

func parseCmdAction(argv []string) EtcdActionInterface {
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

func parseCmdGetAction(argv []string) EtcdActionInterface {
	argc := len(argv)

	switch {
	case argc < argc2:
		return EtcdActionGet{EtcdActGet, "", ""}
	case argc == argc2:
		return EtcdActionGet{EtcdActGet, argv[1], ""}
	case argc > argc2:
		return EtcdActionGet{EtcdActGet, argv[1], argv[2]}
	default:
		return EtcdActionGet{EtcdActGet, "", ""}
	}
}

func parseCmdPutAction(argv []string) EtcdActionInterface {
	if argc := len(argv); argc < argc3 {
		return EtcdActionPut{EtcdActPut, "", ""}
	}

	return EtcdActionPut{EtcdActPut, argv[1], argv[2]}
}

func parseCmdDeleteAction(argv []string) EtcdActionInterface {
	argc := len(argv)

	switch {
	case argc < argc2:
		return EtcdActionDelete{EtcdActDelete, "", ""}
	case argc == argc2:
		return EtcdActionDelete{EtcdActDelete, argv[1], ""}
	case argc > argc2:
		return EtcdActionDelete{EtcdActDelete, argv[1], argv[2]}
	default:
		return EtcdActionDelete{EtcdActDelete, "", ""}
	}
}
