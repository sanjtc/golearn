package cmdclient

import (
	"testing"

	"github.com/pantskun/golearn/customEtcdclt/etcdinteraction"
)

func TestParseCmdAction(t *testing.T) {
	type cmdTestInfo struct {
		cmd      []string
		expected etcdinteraction.EtcdActionInterface
	}

	noneCmd := make([]string, 0)
	cmdCases := []cmdTestInfo{
		{noneCmd, nil},
		{[]string{"other"}, nil},

		{[]string{"get"}, etcdinteraction.NewGetAction("", "")},
		{[]string{"get", "key"}, etcdinteraction.NewGetAction("key", "")},
		{[]string{"get", "key", "rangeEnd"}, etcdinteraction.NewGetAction("key", "rangeEnd")},
		{[]string{"get", "key", "rangeEnd", "other"}, etcdinteraction.NewGetAction("key", "rangeEnd")},

		{[]string{"put"}, etcdinteraction.NewPutAction("", "")},
		{[]string{"put", "key"}, etcdinteraction.NewPutAction("key", "")},
		{[]string{"put", "key", "value"}, etcdinteraction.NewPutAction("key", "value")},
		{[]string{"put", "key", "value", "other"}, etcdinteraction.NewPutAction("key", "value")},

		{[]string{"del"}, etcdinteraction.NewDeleteAction("", "")},
		{[]string{"del", "key"}, etcdinteraction.NewDeleteAction("key", "")},
		{[]string{"del", "key", "rangeEnd"}, etcdinteraction.NewDeleteAction("key", "rangeEnd")},
		{[]string{"del", "key", "rangeEnd", "other"}, etcdinteraction.NewDeleteAction("key", "rangeEnd")},
	}

	for _, cmdCase := range cmdCases {
		action := parseCmdAction(cmdCase.cmd)
		if action == nil && action != cmdCase.expected {
			t.Error("get unexpect action")
		}

		if action != nil && !action.Equal(cmdCase.expected) {
			t.Error("expected: ", cmdCase.expected, " got: ", action)
		}
	}
}
