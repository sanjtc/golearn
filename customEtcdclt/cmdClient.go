package main

import "fmt"

// ParseCommand parse command and execute
func ParseCommand(argv []string) {
	action := parseCmdAction(argv)
	if action == nil {
		return
	}

	msgs, err := action.Exec()

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, msg := range msgs {
		fmt.Println(msg)
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
	if argc := len(argv); argc < 2 {
		return EtcdActionGet{&EtcdActionBase{EtcdActGet}, "", ""}
	} else if argc > 2 {
		return EtcdActionGet{&EtcdActionBase{EtcdActGet}, argv[1], argv[2]}
	} else {
		return EtcdActionGet{&EtcdActionBase{EtcdActGet}, argv[1], ""}
	}
}

func parseCmdPutAction(argv []string) EtcdActionInterface {
	if argc := len(argv); argc < 3 {
		return EtcdActionPut{&EtcdActionBase{EtcdActPut}, "", ""}
	}
	return EtcdActionPut{&EtcdActionBase{EtcdActPut}, argv[1], argv[2]}
}

func parseCmdDeleteAction(argv []string) EtcdActionInterface {
	if argc := len(argv); argc < 2 {
		return EtcdActionDelete{&EtcdActionBase{EtcdActDelete}, "", ""}
	} else if argc > 2 {
		return EtcdActionDelete{&EtcdActionBase{EtcdActDelete}, argv[1], argv[2]}
	} else {
		return EtcdActionDelete{&EtcdActionBase{EtcdActDelete}, argv[1], ""}
	}
}
