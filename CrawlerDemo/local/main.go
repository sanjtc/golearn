package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/pantskun/commonutils/osutils"
	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

const interruptMsg = "interrupt"

func main() {
	// 计时
	startTime := time.Now()

	defer func() {
		useTime := time.Since(startTime)
		xlogutil.Warning("use time:", useTime)
	}()

	interruptChan := make(chan int)
	// 处理远程中断
	processRemoteInterrupt(":2233", interruptChan)
	// 处理本地中断
	processLocalInterrupt(interruptChan)

	// cmd参数
	var (
		procNum int
		url     string
	)

	flag.IntVar(&procNum, "n", 1, "process number")
	flag.StringVar(&url, "url", "https://www.ssetech.com.cn/", "url")
	flag.Parse()

	// 启动etcd
	startEtcdCmd := osutils.NewCommand("etcd")
	startEtcdCmd.RunAsyn()

	defer func() {
		// 清除etcd数据，关闭etcd
		delEtcdDataCmd := osutils.NewCommand("etcdctl", "del", "--prefix", "/crawler")
		delEtcdDataCmd.Run()

		_ = startEtcdCmd.Kill()
	}()

	waitChan := make(chan int)

	// 开启n个crawler进程
	modulePath := pathutils.GetModulePath("CrawlerDemo")
	mainPath := path.Join(modulePath, "crawler", "main.go")

	cmds := make([]osutils.Command, procNum)
	for i := 0; i < procNum; i++ {
		cmds[i] = osutils.NewCommand(
			"go",
			"run", mainPath,
			"-url", url,
			"-depth", "2",
			"-log", "crawler"+strconv.Itoa(i)+".txt",
			"-useMultiprocess",
		)
		cmds[i].RunAsyn()
	}

	// 等待所有crawler进程执行完成
	go func() {
		for _, cmd := range cmds {
			for cmd.GetCmdState() == osutils.ECmdStateRunning {
			}
		}
		waitChan <- 0
	}()

	// 等待结果
	select {
	case <-waitChan:
		{
			// 日志文件路径
			logPaths := []string{}
			for i := 0; i < procNum; i++ {
				logPaths = append(logPaths, path.Join(modulePath, "logs", "crawler"+strconv.Itoa(i)+".txt"))
			}

			// 读取日志内容
			fileContents := loadLogs(logPaths)

			// 检查执行结果
			if checkLogs(fileContents) {
				xlogutil.Warning("successed")
			} else {
				xlogutil.Warning("failed")
			}
		}
	case <-interruptChan:
		{
			// 结束所有crawler进程
			for _, cmd := range cmds {
				if cmd == nil {
					continue
				}
				if err := cmd.Kill(); err != nil {
					xlogutil.Error(err)
				}
			}
			xlogutil.Warning(interruptMsg)
		}
	}
}

func loadLogs(filePaths []string) (fileContents [][]string) {
	n := len(filePaths)
	fileContents = make([][]string, n)

	for i, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			xlogutil.Error(err)
			return nil
		}

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			xlogutil.Error(err)
			return nil
		}

		fileContents[i] = strings.Split(string(fileContent), "\n")
	}

	return
}

func checkLogs(fileContents [][]string) bool {
	logs := make(map[string]int)

	for i, fileContent := range fileContents {
		if len(fileContent) == 0 {
			return false
		}

		if fileContent[len(fileContent)-1] == "" {
			fileContent = fileContent[:len(fileContent)-1]
		}

		if !strings.Contains(fileContent[len(fileContent)-1], "finished") {
			xlogutil.Error("crawler " + strconv.Itoa(i) + ": not finished")
			return false
		}

		for _, log := range fileContent {
			if !strings.Contains(log, "download:") && !strings.Contains(log, "process:") {
				continue
			}

			if logs[log] == 0 {
				logs[log] = 1
				continue
			}

			if logs[log] == 1 {
				xlogutil.Error("crawler " + strconv.Itoa(i) + ": download or process duplicate content")
				xlogutil.Error(log)

				return false
			}
		}
	}

	return true
}

func processLocalInterrupt(interruptChan chan int) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	signalChan := make(chan os.Signal)
	osutils.ListenSystemSignalsWithCtx(ctx, cancel, signalChan, os.Interrupt)

	go func() {
		<-ctx.Done()
		interruptChan <- 0
	}()
}

func processRemoteInterrupt(listenAddr string, interruptChan chan int) {
	go func() {
		err := listenRemoteInterrupt(listenAddr, interruptChan)
		if err != nil {
			xlogutil.Error(err)
		}
	}()
}

func listenRemoteInterrupt(addr string, interruptChan chan int) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		buf := make([]byte, 20)
		size, err := conn.Read(buf)
		conn.Close()

		if err != nil {
			return err
		}

		msg := string(buf[:size])

		xlogutil.Warning("receive:", msg)

		if msg == interruptMsg {
			// 处理中断
			interruptChan <- 0
			return nil
		}
	}
}
