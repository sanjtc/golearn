package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/pantskun/commonutils/osutils"
	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/commonutils/taskutils"
	"github.com/pantskun/remotelib/remotesftp"
	"github.com/pantskun/remotelib/remotessh"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal)
	osutils.ListenSystemSignalsWithCtx(ctx, cancel, signalChan, os.Interrupt)

	taskPool := taskutils.NewTaskPool()
	taskPool.Run()

	uploadTask := taskutils.NewTask(
		"uploadTask",
		func() error {
			log.Println("uploadTask start")

			if err := UploadSrc(); err != nil {
				return err
			}

			log.Println("uploadTask finished")
			return nil
		},
	)
	taskPool.AddTask(uploadTask)

	runTask := taskutils.NewTask(
		"runTask",
		func() error {
			log.Println("runTask start")

			if err := RunSrc(); err != nil {
				return err
			}

			log.Println("runTask finished")
			return nil
		},
		uploadTask,
	)
	taskPool.AddTask(runTask)

	// 等待任务完成，或者检测到中断发送中断到远程任务
	for {
		if uploadTask.GetState() == taskutils.ETaskStateFinished &&
			runTask.GetState() == taskutils.ETaskStateFinished {
			break
		}

		select {
		case <-ctx.Done():
			{
				if err := ProcessInterrupt(); err != nil {
					log.Println(err)
				}

				break
			}
		default:
			{
				continue
			}
		}
	}

	taskPool.Close()
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
		"cd /home/wx/CrawlerDemo/localScript",
		"go run main.go -n 4",
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
