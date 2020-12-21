package main

import (
	"context"
	"net"
	"path"
	"testing"
	"time"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/stretchr/testify/assert"
)

func TestProcessRemoteInterrupt(t *testing.T) {
	interruptChan := make(chan int, 1)

	go processRemoteInterrupt("localhost:2333", interruptChan)

	conn, err := net.Dial("tcp", "localhost:2333")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	msg := "interrupt"
	data := []byte(msg)

	_, err = conn.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	receviced := false

	go func() {
		<-interruptChan

		receviced = true

		cancel()
	}()

	<-ctx.Done()

	assert.True(t, receviced)
}

func TestCheckLogs(t *testing.T) {
	type TestCase struct {
		fileContent [][]string
		expected    bool
	}

	testCases := []TestCase{
		{
			fileContent: [][]string{
				{"download:test0", "finished"},
				{"download:test1", "finished"},
			},
			expected: true,
		},
		{
			fileContent: [][]string{
				{"download:test0", "finished"},
				{"download:test0", "finished"},
			},
			expected: false,
		},
		{
			fileContent: [][]string{
				{"download:test0"},
				{"download:test1", "finished"},
			},
			expected: false,
		},
		{
			fileContent: [][]string{
				{},
				{"download:test1", "finished"},
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		got := checkLogs(testCase.fileContent)
		assert.Equal(t, testCase.expected, got)
	}
}

func TestLoadLogs(t *testing.T) {
	type TestCase struct {
		filePaths []string
		expected  [][]string
	}

	modulePath := pathutils.GetModulePath("CrawlerDemo")

	testCases := []TestCase{
		{
			filePaths: []string{
				path.Join(modulePath, "logs", "test", "test0.txt"),
				path.Join(modulePath, "logs", "test", "test1.txt"),
			},
			expected: [][]string{
				{"test0"},
				{"test1"},
			},
		},
		{
			filePaths: []string{
				path.Join(modulePath, "logs", "test", "test888.txt"),
				path.Join(modulePath, "logs", "test", "test1.txt"),
			},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		got := loadLogs(testCase.filePaths)
		assert.True(t, logContentEqual(testCase.expected, got))
	}
}

func logContentEqual(expected, got [][]string) bool {
	n := len(expected)
	if n != len(got) {
		return false
	}

	for i := 0; i < n; i++ {
		elog := expected[i]
		glog := got[i]

		m := len(elog)
		if m != len(glog) {
			return false
		}

		for j := 0; j < m; j++ {
			if elog[j] != glog[j] {
				return false
			}
		}
	}

	return true
}
