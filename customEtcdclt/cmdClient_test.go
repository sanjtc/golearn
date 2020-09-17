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
	noneCmd := make([]string, 0)
	cmdCases := []cmdTestInfo{
		{noneCmd, nil},
		{[]string{"other"}, nil},

		{[]string{"get"}, EtcdActionGet{EtcdActionBase{EtcdActGet}, "", ""}},
		{[]string{"get", "key"}, EtcdActionGet{EtcdActionBase{EtcdActGet}, "key", ""}},
		{[]string{"get", "key", "rangeEnd"}, EtcdActionGet{EtcdActionBase{EtcdActGet}, "key", "rangeEnd"}},
		{[]string{"get", "key", "rangeEnd", "other"}, EtcdActionGet{EtcdActionBase{EtcdActGet}, "key", "rangeEnd"}},

		{[]string{"put"}, EtcdActionPut{EtcdActionBase{EtcdActPut}, "", ""}},
		{[]string{"put", "key"}, EtcdActionPut{EtcdActionBase{EtcdActPut}, "key", ""}},
		{[]string{"put", "key", "value"}, EtcdActionPut{EtcdActionBase{EtcdActPut}, "key", "value"}},
		{[]string{"put", "key", "value", "other"}, EtcdActionPut{EtcdActionBase{EtcdActPut}, "key", "value"}},

		{[]string{"del"}, EtcdActionDelete{EtcdActionBase{EtcdActDelete}, "", ""}},
		{[]string{"del", "key"}, EtcdActionDelete{EtcdActionBase{EtcdActDelete}, "key", ""}},
		{[]string{"del", "key", "rangeEnd"}, EtcdActionDelete{EtcdActionBase{EtcdActDelete}, "key", "rangeEnd"}},
		{[]string{"del", "key", "rangeEnd", "other"}, EtcdActionDelete{EtcdActionBase{EtcdActDelete}, "key", "rangeEnd"}},
	}

	for _, cmdCase := range cmdCases {
		action := parseCmdAction(cmdCase.cmd)
		if action == nil && action != cmdCase.action {
			t.Errorf("get unexpect action")
		}
		if action != nil && !action.Equal(cmdCase.action) {
			fmt.Println(action)
			fmt.Println(cmdCase.action)
			t.Errorf("get unexpect action")
		}
	}
}
