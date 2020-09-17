package main

import (
	"fmt"
	"testing"
)

type cmdTestInfo struct {
	cmd    []string
	action EtcdActionInterface
}

func TestParseCmdAction(t *testing.T) {

	cmdCases := []cmdTestInfo{
		{nil, nil},
		{[]string{"other"}, nil},
		{[]string{"get"}, EtcdActionGet{&EtcdActionBase{EtcdActGet}, "", ""}},
		{[]string{"get", "key"}, EtcdActionGet{&EtcdActionBase{EtcdActGet}, "key", ""}},
		{[]string{"get", "key", "rangeEnd"}, EtcdActionGet{&EtcdActionBase{EtcdActGet}, "key", "rangeEnd"}},
		{[]string{"get", "key", "rangeEnd", "other"}, EtcdActionGet{&EtcdActionBase{EtcdActGet}, "key", "rangeEnd"}},
	}

	for _, cmdCase := range cmdCases {
		action := parseCmdAction(cmdCase.cmd)
		if action.Equal(cmdCase.action) {
			fmt.Println(action)
			fmt.Println(cmdCase.action)
			t.Errorf("get unexpect action")
		}
	}

	invalidCmd := [][]string{
		nil,
		[]string{"other"},
	}
	for _, cmd := range invalidCmd {
		action := parseCmdAction(cmd)
		if action != nil {
			t.Errorf("none cmd can not get action")
		}
	}

	getCmds := [][]string{
		[]string{"get"},
		[]string{"get", "key"},
		[]string{"get", "key", "rangeEnd"},
		[]string{"get", "key", "rangeEnd", "other"},
	}
	for _, getCmd := range getCmds {
		action := parseCmdAction(getCmd)
		if _, ok := action.(EtcdActionGet); !ok {
			t.Errorf("getCmd can not get getAction")
		}
	}

	putCmds := [][]string{
		[]string{"put"},
		[]string{"put", "key"},
		[]string{"put", "key", "value"},
		[]string{"put", "key", "value", "other"},
	}
	for _, putCmd := range putCmds {
		action := parseCmdAction(putCmd)
		if _, ok := action.(EtcdActionPut); !ok {
			t.Errorf("putCmd can not return put action")
		}
	}

	deleteCmds := [][]string{
		[]string{"del"},
		[]string{"del", "key"},
		[]string{"del", "key", "rangeEnd"},
		[]string{"del", "key", "rangeEnd", "other"},
	}
	for _, deleteCmd := range deleteCmds {
		action := parseCmdAction(deleteCmd)
		if _, ok := action.(EtcdActionDelete); !ok {
			t.Errorf("deleteCmd can not return delete action")
		}
	}
}
