package main

import (
	"net"
	"testing"
	"time"

	"github.com/pantskun/commonutils/pathutils"
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

	err = sftpInteractor.Upload(pathutils.GetModulePath("CrawlerDemo"), "/home/wx")
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

	cmds := []string{
		"cd /home/wx/CrawlerDemo",
		"go run start.go -n 4",
	}

	out, err := sshInteractor.Run(cmds)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(out)
}

func TestChan(t *testing.T) {
	ch := make(chan struct{})

	wait := func() <-chan struct{} {
		return ch
	}

	go func() {
		<-wait()
		t.Log("success")
	}()

	time.Sleep(5 * time.Second)

	ch <- struct{}{}
}

func TestListenRemoteInterrupt(t *testing.T) {
	ch := make(chan int)
	if err := listenRemoteInterrupt(":2233", ch); err != nil {
		t.Fatal(err)
	}
}

func TestAddr(t *testing.T) {
	t.Log(net.Interfaces())
}
