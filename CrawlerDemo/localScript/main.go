package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path"
	"strings"

	"github.com/pantskun/commonutils/osutils"
	"github.com/pantskun/commonutils/pathutils"
)

func main() {
	var procNum int

	flag.IntVar(&procNum, "n", 1, "process number")
	flag.Parse()

	// 启动etcd
	startEtcdCmd := osutils.NewCommand("etcd")

	go func() {
		startEtcdCmd.Run()

		if startEtcdCmd.GetCmdState() == osutils.ECmdStateError {
			log.Println(startEtcdCmd.GetCmdError())
			return
		}

		log.Println(startEtcdCmd.GetStdout())
	}()

	defer func() {
		// 清除etcd数据，关闭etcd
		delEtcdDataCmd := osutils.NewCommand("etcdctl", "del", "--prefix", "https://")
		delEtcdDataCmd.Run()

		if delEtcdDataCmd.GetCmdState() == osutils.ECmdStateError {
			log.Println(delEtcdDataCmd.GetCmdError())
		}

		_ = startEtcdCmd.Kill()
	}()

	// start n processes
	mainPath := path.Join(pathutils.GetModulePath("CrawlerDemo"), "main", "main.go")
	log.Println("work directory: ", mainPath)

	cmds := make([]osutils.Command, procNum)

	for i := 0; i < procNum; i++ {
		cmds[i] = osutils.NewCommand("go", "run", mainPath)
		cmds[i].RunAsyn()
		log.Println("start process ", i+1)
	}

	// 是否需要对多个进程执行结果进行检查
	var needCheck bool = true

	// 等待多个进程执行完成
	waitChan := make(chan int)

	waitProc := func() {
		for _, cmd := range cmds {
			for cmd.GetCmdState() == osutils.ECmdStateRunning {
			}

			// 如果有cmd执行失败，则不进行check
			if cmd.GetCmdState() == osutils.ECmdStateError {
				needCheck = false
			}
		}
		waitChan <- 0
	}
	go waitProc()

	// 处理远程中断
	interruptChan := make(chan int)

	go func() {
		err := listenRemoteInterrupt(":2233", interruptChan)
		if err != nil {
			log.Println(err)
		}
	}()

	select {
	case <-waitChan:
		{
			// 检查执行结果
			if needCheck && checkCmds(cmds) {
				log.Println("successed")
			} else {
				log.Println("failed")
			}
		}
	case <-interruptChan:
		{
			log.Println("receive interrupt")

			for _, cmd := range cmds {
				if cmd == nil {
					continue
				}
				if err := cmd. /*Process.*/ Kill(); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func checkCmds(cmds []osutils.Command) bool {
	n := len(cmds)
	outs := make([][]string, n)

	for i := 0; i < n; i++ {
		outs[i] = strings.Split(cmds[i].GetStdout(), "\n")
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

		if msg == "interrupt" {
			// 处理中断
			interruptChan <- 0
			return nil
		}
	}
}
