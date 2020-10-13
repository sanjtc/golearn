package main

import (
	"testing"

	"github.com/pantskun/golearn/CrawlerDemo/pathutils"
)

func TestMain(t *testing.T) {
	t.Log(pathutils.GetModulePath())
}
