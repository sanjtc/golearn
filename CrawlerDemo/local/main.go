package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path"
	"runtime"
	"strings"

	"github.com/pantskun/commonutils/osutils"
	"github.com/pantskun/commonutils/pathutils"
)

const interruptMsg = "interrupt"

func main() {
	var (
		procNum int
		url     string
	)

	flag.IntVar(&procNum, "n", 1, "process number")
	flag.StringVar(&url, "url", "https://www.ssetech.com.cn/", "url")
	flag.Parse()

	if procNum > runtime.NumCPU() {
		procNum = runtime.NumCPU()
		log.Println("Number of CPU core is ", runtime.NumCPU())
	}

	// 启动etcd
	startEtcdCmd := osutils.NewCommand("etcd")
	startEtcdCmd.RunAsyn()

	defer func() {
		// 清除etcd数据，关闭etcd
		delEtcdDataCmd := osutils.NewCommand("etcdctl", "del", "--prefix", "https://")
		delEtcdDataCmd.Run()

		_ = startEtcdCmd.Kill()
	}()

	waitChan := make(chan int)
	interruptChan := make(chan int)

	// start n processes
	mainPath := path.Join(pathutils.GetModulePath("CrawlerDemo"), "crawler", "main.go")
	multiProcCmd := osutils.NewMultiProcCmd(procNum, "go", "run", mainPath, "-url", url)

	if startEtcdCmd.GetCmdState() == osutils.ECmdStateError {
		log.Println(startEtcdCmd.GetCmdError())
		return
	}

	go func() {
		multiProcCmd.Run()
		waitChan <- 0
	}()

	// 处理远程中断
	processRemoteInterrupt(":2233", interruptChan)

	// 等待结果
	log.Println(waitingResult(waitChan, interruptChan, multiProcCmd.GetCmds()))
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
					log.Println(err)
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
			log.Println(err)
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

		log.Println("receive:", msg)

		if msg == interruptMsg {
			// 处理中断
			interruptChan <- 0
			return nil
		}
	}
}
