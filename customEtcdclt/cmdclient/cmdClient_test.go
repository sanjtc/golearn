package cmdclient

import (
	"testing"

	"github.com/pantskun/golearn/customEtcdclt/etcdinteraction"
)

func TestParseCmdAction(t *testing.T) {
	type cmdTestInfo struct {
		cmd    []string
		action etcdinteraction.EtcdActionInterface
	}

	noneCmd := make([]string, 0)
	cmdCases := []cmdTestInfo{
		{noneCmd, nil},
		{[]string{"other"}, nil},

		{[]string{"get"}, etcdinteraction.EtcdActionGet{ActionType: etcdinteraction.EtcdActGet, Key: "", RangeEnd: ""}},
		{[]string{"get", "key"}, etcdinteraction.EtcdActionGet{ActionType: etcdinteraction.EtcdActGet, Key: "key", RangeEnd: ""}},
		{[]string{"get", "key", "rangeEnd"}, etcdinteraction.EtcdActionGet{ActionType: etcdinteraction.EtcdActGet, Key: "key", RangeEnd: "rangeEnd"}},
		{[]string{"get", "key", "rangeEnd", "other"}, etcdinteraction.EtcdActionGet{ActionType: etcdinteraction.EtcdActGet, Key: "key", RangeEnd: "rangeEnd"}},

		{[]string{"put"}, etcdinteraction.EtcdActionPut{ActionType: etcdinteraction.EtcdActPut, Key: "", Value: ""}},
		{[]string{"put", "key"}, etcdinteraction.EtcdActionPut{ActionType: etcdinteraction.EtcdActPut, Key: "", Value: ""}},
		{[]string{"put", "key", "value"}, etcdinteraction.EtcdActionPut{ActionType: etcdinteraction.EtcdActPut, Key: "key", Value: "value"}},
		{[]string{"put", "key", "value", "other"}, etcdinteraction.EtcdActionPut{ActionType: etcdinteraction.EtcdActPut, Key: "key", Value: "value"}},

		{[]string{"del"}, etcdinteraction.EtcdActionDelete{ActionType: etcdinteraction.EtcdActDelete, Key: "", RangeEnd: ""}},
		{[]string{"del", "key"}, etcdinteraction.EtcdActionDelete{ActionType: etcdinteraction.EtcdActDelete, Key: "key", RangeEnd: ""}},
		{[]string{"del", "key", "rangeEnd"}, etcdinteraction.EtcdActionDelete{ActionType: etcdinteraction.EtcdActDelete, Key: "key", RangeEnd: "rangeEnd"}},
		{[]string{"del", "key", "rangeEnd", "other"}, etcdinteraction.EtcdActionDelete{ActionType: etcdinteraction.EtcdActDelete, Key: "key", RangeEnd: "rangeEnd"}},
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
