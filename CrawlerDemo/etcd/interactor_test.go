package etcd

import (
	"testing"

	"github.com/pantskun/commonutils/osutils"
	"github.com/stretchr/testify/assert"
)

func TestInteractorError(t *testing.T) {
	type errorCase struct {
		msg      string
		expected string
	}

	testCases := []errorCase{
		{msg: "msg", expected: "msg"},
	}

	for _, testCase := range testCases {
		err := InteractorError{testCase.msg}
		got := err.Error()
		assert.Equal(t, testCase.expected, got)
	}
}

func TestNewInteractorWithEmbed(t *testing.T) {
	interactor, err := NewInteractorWithEmbed()
	if err != nil {
		t.Log(err)
		return
	}
	defer interactor.Close()

	testPut(t, interactor)
	testGet(t, interactor)
	testDel(t, interactor)
	testLock(t, interactor)
}

func testPut(t *testing.T, interactor Interactor) {
	type putCase struct {
		key      string
		value    string
		expected error
	}

	testCases := []putCase{
		{key: "key1", value: "value1", expected: nil},
	}

	for _, testCase := range testCases {
		got := interactor.Put(testCase.key, testCase.value)
		assert.Equal(t, testCase.expected, got)
	}
}

func testGet(t *testing.T, interactor Interactor) {
	type getCase struct {
		key      string
		expected string
	}

	testCases := []getCase{
		{key: "key1", expected: "value1"},
		{key: "key2", expected: ""},
	}

	for _, testCase := range testCases {
		got, err := interactor.Get(testCase.key)
		if err != nil {
			t.Fatal()
		}

		assert.Equal(t, testCase.expected, got)
	}
}

func testDel(t *testing.T, interactor Interactor) {
	type delCase struct {
		key      string
		expected error
	}

	testCases := []delCase{
		{key: "key1", expected: nil},
	}

	for _, testCase := range testCases {
		got := interactor.Del(testCase.key)
		assert.Equal(t, testCase.expected, got)
	}
}

func testLock(t *testing.T, interactor Interactor) {
	if _, err := interactor.Lock(); err != nil {
		t.Log(err)
	}

	if _, err := interactor.Unlock(); err != nil {
		t.Log(err)
	}
}

func TestNewInteractor(t *testing.T) {
	etcdCmd := osutils.NewCommand("etcd")
	etcdCmd.RunAsyn()

	defer etcdCmd.Kill()

	interactor, err := NewInteractor()
	if err != nil {
		t.Log(err)
	} else {
		interactor.Close()
	}
}

func TestTxn(t *testing.T) {
	etcdEmbed, _ := NewInteractorWithEmbed()
	assert.True(t, etcdEmbed.TxnSync("test1"))
	assert.False(t, etcdEmbed.TxnSync("test1"))
}
