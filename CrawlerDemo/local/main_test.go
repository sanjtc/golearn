package main

import (
	"path"
	"testing"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/stretchr/testify/assert"
)

func TestCheckLogs(t *testing.T) {
	type TestCase struct {
		fileContent [][]string
		expected    bool
	}

	testCases := []TestCase{
		{
			fileContent: [][]string{
				[]string{"download:test0", "finished"},
				[]string{"download:test1", "finished"},
			},
			expected: true,
		},
		{
			fileContent: [][]string{
				[]string{"download:test0", "finished"},
				[]string{"download:test0", "finished"},
			},
			expected: false,
		},
		{
			fileContent: [][]string{
				[]string{"download:test0"},
				[]string{"download:test1", "finished"},
			},
			expected: false,
		},
		{
			fileContent: [][]string{
				[]string{},
				[]string{"download:test1", "finished"},
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
				[]string{"test0"},
				[]string{"test1"},
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
