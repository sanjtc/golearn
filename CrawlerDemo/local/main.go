package main

import (
	"flag"
	"fmt"
	"net"
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
	startTime := time.Now()

	defer func() {
		useTime := time.Since(startTime)
		xlogutil.Warning("use time:", useTime)
	}()

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
	interruptChan := make(chan int)

	// start n processes
	mainPath := path.Join(pathutils.GetModulePath("CrawlerDemo"), "crawler", "main.go")

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

	go func() {
		for _, cmd := range cmds {
			for cmd.GetCmdState() == osutils.ECmdStateRunning {
			}
		}
		waitChan <- 0
	}()

	// 处理远程中断
	processRemoteInterrupt(":2233", interruptChan)

	// 等待结果
	xlogutil.Warning(waitingResult(waitChan, interruptChan, cmds /*multiProcCmd.GetCmds()*/))
}

func waitingResult(waitChan, interruptChan chan int, cmds []osutils.Command) string {
	select {
	case <-waitChan:
		{
			// 检查执行结果
			if checkCmdOuts(cmds) {
				return "successed"
			} else {
				return "failed"
			}
		}
	case <-interruptChan:
		{
			for _, cmd := range cmds {
				if cmd == nil {
					continue
				}
				if err := cmd. /*Process.*/ Kill(); err != nil {
					xlogutil.Error(err)
				}
			}
			return interruptMsg
		}
	}
}

func checkCmdOuts(cmds []osutils.Command) bool {
	n := len(cmds)
	outs := make([][]string, n)

	for i, cmd := range cmds {
		// 有命令执行失败，返回false
		if cmd.GetCmdState() == osutils.ECmdStateError {
			return false
		}

		outs[i] = strings.Split(cmd.GetStdout(), "\n")
		outs[i] = outs[i][0 : len(outs[i])-1]
	}

	maps := make(map[string]int)

	for _, out := range outs {
		for _, url := range out {
			if maps[url] == 0 {
				maps[url] = 1
				continue
			}

			if maps[url] == 1 {
				fmt.Println(url)
				return false
			}
		}
	}

	return true
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
