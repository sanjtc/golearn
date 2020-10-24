package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/pantskun/commonutils/osutils"
	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/remotelib/remotesftp"
	"github.com/pantskun/remotelib/remotessh"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal)
	osutils.ListenSystemSignalsWithCtx(ctx, cancel, signalChan, os.Interrupt)

	taskChan := make(chan int)

	go func() {
		if err := UploadSrc(); err != nil {
			log.Println(err)
			return
		}

		if err := RunSrc(); err != nil {
			log.Println(err)
			return
		}

		taskChan <- 0
	}()

	select {
	case <-taskChan:
		{
			return
		}
	case <-ctx.Done():
		{
			if err := ProcessInterrupt(); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func UploadSrc() error {
	sftpConfig := remotesftp.SFTPConfig{
		Network:  "tcp",
		IP:       "192.168.62.11",
		Port:     "22",
		User:     "wx",
		Password: "1235",
	}

	sftpInteractor, err := remotesftp.NewInteractor(sftpConfig)
	if err != nil {
		return err
	}

	err = sftpInteractor.Upload(pathutils.GetModulePath("CrawlerDemo"), "/home/wx")
	if err != nil {
		return err
	}

	return nil
}

func RunSrc() error {
	sshConfig := remotessh.SSHConfig{
		Network:  "tcp",
		IP:       "192.168.62.11",
		Port:     "22",
		User:     "wx",
		Password: "1235",
	}

	sshInteractor, err := remotessh.NewInteractor(sshConfig)
	if err != nil {
		return err
	}

	cmds := []string{
		"cd /home/wx/CrawlerDemo/local",
		"go run start.go -n 4",
	}

	out, err := sshInteractor.Run(cmds)
	if err != nil {
		return err
	}

	log.Println(out)

	return nil
}

func ProcessInterrupt() error {
	conn, err := net.Dial("tcp", "192.168.62.11:2233")
	if err != nil {
		return err
	}
	defer conn.Close()

	msg := "interrupt"
	data := []byte(msg)

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}
