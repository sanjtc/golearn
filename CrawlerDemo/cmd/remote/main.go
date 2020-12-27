package main

import (
	"context"
	"flag"
	"net"
	"os"
	"strconv"

	"github.com/pantskun/commonutils/osutils"
	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/commonutils/taskutils"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
	"github.com/pantskun/remotelib/remotesftp"
	"github.com/pantskun/remotelib/remotessh"
)

func main() {
	var (
		remoteIP   string
		remotePort string
		remoteUser string
		remotePwd  string
		uploadPath string

		procNum int

		url string
	)

	flag.StringVar(&remoteIP, "addr", "192.168.62.11", "remote address")
	flag.StringVar(&remotePort, "port", "22", "remote port")
	flag.StringVar(&remoteUser, "user", "wx", "remote user")
	flag.StringVar(&remotePwd, "pwd", "1235", "remote password")
	flag.StringVar(&uploadPath, "path", "/home/wx", "upload path")
	flag.IntVar(&procNum, "n", 4, "process number")
	flag.StringVar(&url, "url", "https://www.ssetech.com.cn/", "url")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	signalChan := make(chan os.Signal)
	osutils.ListenSystemSignalsWithCtx(ctx, cancel, signalChan, os.Interrupt)

	taskPool := taskutils.NewTaskPool()
	taskPool.Run()

	defer taskPool.Close()

	uploadTask := taskutils.NewTask(
		"uploadTask",
		func() error {
			xlogutil.Warning("uploadTask start")

			if err := uploadSrc(remoteIP, remotePort, remoteUser, remotePwd, uploadPath); err != nil {
				return err
			}

			xlogutil.Warning("uploadTask finished")
			return nil
		},
	)
	taskPool.AddTask(uploadTask)

	runTask := taskutils.NewTask(
		"runTask",
		func() error {
			xlogutil.Warning("runTask start")

			if err := runSrc(remoteIP, remotePort, remoteUser, remotePwd, procNum, url); err != nil {
				return err
			}

			xlogutil.Warning("runTask finished")
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

		if uploadTask.GetState() == taskutils.ETaskStateError {
			xlogutil.Warning("uploadTask execute failed")
			return
		}

		if runTask.GetState() == taskutils.ETaskStateError {
			xlogutil.Warning("runTask execute failed")
			return
		}

		select {
		case <-ctx.Done():
			{
				if err := processInterrupt(remoteIP, "2233"); err != nil {
					xlogutil.Error(err)
				}
			}
		default:
			{
				continue
			}
		}
	}
}

func uploadSrc(ip, port, user, pwd, uploadPath string) error {
	sftpConfig := remotesftp.SFTPConfig{
		Network:  "tcp",
		IP:       ip,
		Port:     port,
		User:     user,
		Password: pwd,
	}

	sftpInteractor, err := remotesftp.NewInteractor(sftpConfig)
	if err != nil {
		return err
	}

	err = sftpInteractor.Upload(pathutils.GetModulePath("CrawlerDemo"), uploadPath)
	if err != nil {
		return err
	}

	return nil
}

func runSrc(ip, port, user, pwd string, procNum int, url string) error {
	sshConfig := remotessh.SSHConfig{
		Network:  "tcp",
		IP:       ip,
		Port:     port,
		User:     user,
		Password: pwd,
	}

	sshInteractor, err := remotessh.NewInteractor(sshConfig)
	if err != nil {
		return err
	}

	cmds := []string{
		"cd /home/wx/CrawlerDemo/cmd/local",
		"go run main.go -n " + strconv.Itoa(procNum) + " -url " + url,
	}

	if err := sshInteractor.Run(cmds); err != nil {
		return err
	}

	xlogutil.Warning("stderr: \n", sshInteractor.GetStderr())

	return nil
}

func processInterrupt(ip, port string) error {
	conn, err := net.Dial("tcp", ip+":"+port)
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
