package main

import (
	"testing"
)

func TestParseCmdAction(t *testing.T) {
	type cmdTestInfo struct {
		cmd    []string
		action EtcdActionInterface
	}

	noneCmd := make([]string, 0)
	cmdCases := []cmdTestInfo{
		{noneCmd, nil},
		{[]string{"other"}, nil},

		{[]string{"get"}, EtcdActionGet{EtcdActGet, "", ""}},
		{[]string{"get", "key"}, EtcdActionGet{EtcdActGet, "key", ""}},
		{[]string{"get", "key", "rangeEnd"}, EtcdActionGet{EtcdActGet, "key", "rangeEnd"}},
		{[]string{"get", "key", "rangeEnd", "other"}, EtcdActionGet{EtcdActGet, "key", "rangeEnd"}},

		{[]string{"put"}, EtcdActionPut{EtcdActPut, "", ""}},
		{[]string{"put", "key"}, EtcdActionPut{EtcdActPut, "", ""}},
		{[]string{"put", "key", "value"}, EtcdActionPut{EtcdActPut, "key", "value"}},
		{[]string{"put", "key", "value", "other"}, EtcdActionPut{EtcdActPut, "key", "value"}},

		{[]string{"del"}, EtcdActionDelete{EtcdActDelete, "", ""}},
		{[]string{"del", "key"}, EtcdActionDelete{EtcdActDelete, "key", ""}},
		{[]string{"del", "key", "rangeEnd"}, EtcdActionDelete{EtcdActDelete, "key", "rangeEnd"}},
		{[]string{"del", "key", "rangeEnd", "other"}, EtcdActionDelete{EtcdActDelete, "key", "rangeEnd"}},
	}

	for _, cmdCase := range cmdCases {
		action := parseCmdAction(cmdCase.cmd)
		if action == nil && action != cmdCase.action {
			t.Errorf("get unexpect action")
		}

		if action != nil && !action.Equal(cmdCase.action) {
			t.Errorf("get unexpect action")
		}
	}
}
