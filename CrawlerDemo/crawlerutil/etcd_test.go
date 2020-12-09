package crawlerutil

import (
	"testing"

	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"gotest.tools/assert"
)

func TestSynchronize(t *testing.T) {
	etcdInteractor, err := etcd.NewInteractorWithEmbed()

	if err != nil {
		t.Log(err)
	}

	defer etcdInteractor.Close()

	type TestCase struct {
		key      string
		expected bool
	}

	testCases := []TestCase{
		{"test1", true},
		{"test1", false},
	}

	for _, testCase := range testCases {
		got := Synchronize(testCase.key, etcdInteractor)
		assert.Equal(t, got, testCase.expected)
	}
}
