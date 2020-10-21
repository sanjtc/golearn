package main

import (
	"testing"

	"github.com/pantskun/pathlib"
	"github.com/pantskun/remotelib/remotesftp"
)

func TestStart(t *testing.T) {
	// build main
	// buildCmd := exec.Command("go", "build", "./main/")
	// buildCmd.Dir = pathlib.GetModulePath("CrawlerDemo")

	// err := buildCmd.Run()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	sftpConfig := remotesftp.SFTPConfig{
		Network:  "tcp",
		IP:       "192.168.62.11",
		Port:     "22",
		User:     "wx",
		Password: "1235",
	}

	sftpInteractor, err := remotesftp.NewInteractor(sftpConfig)
	if err != nil {
		t.Fatal(err)
	}

	err = sftpInteractor.Upload(pathlib.GetModulePath("CrawlerDemo"), "/home/wx")
	if err != nil {
		t.Fatal(err)
	}
}
