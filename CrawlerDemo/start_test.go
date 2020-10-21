package main

import (
	"testing"

	"github.com/pantskun/pathlib"
	"github.com/pantskun/remotelib/remotesftp"
	"github.com/pantskun/remotelib/remotessh"
)

func TestUploadSrc(t *testing.T) {
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

func TestRunSrc(t *testing.T) {
	sshConfig := remotessh.SSHConfig{
		Network:  "tcp",
		IP:       "192.168.62.11",
		Port:     "22",
		User:     "wx",
		Password: "1235",
	}

	sshInteractor, err := remotessh.NewInteractor(sshConfig)
	if err != nil {
		t.Fatal(err)
	}

	out, err := sshInteractor.Run("go run /home/wx/CrawlerDemo/start.go")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(out)
}
