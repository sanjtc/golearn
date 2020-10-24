package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"path"
	"strings"

	"github.com/pantskun/commonutils/pathutils"
)

func main() {
	procnum := flag.Int("n", 1, "process number")
	flag.Parse()

	n := *procnum

	// start etcd
	etcdCmd := exec.Command("etcd")

	var etcdOut bytes.Buffer
	etcdCmd.Stdout = &etcdOut

	go func() {
		err := etcdCmd.Run()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(etcdOut)
	}()

	defer func() {
		// clean url data
		etcdctlCmd := exec.Command("etcdctl", "del", "--prefix", "https://")

		var etcdctlOuter bytes.Buffer
		etcdctlCmd.Stdout = &etcdctlOuter

		err := etcdctlCmd.Run()
		if err != nil {
			fmt.Println(err)
		}

		// close etcd
		_ = etcdCmd.Process.Kill()
	}()

	// start n processes
	execCmd := func(cmd *exec.Cmd, s chan int) {
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			s <- 0
		} else {
			s <- 1
		}
	}

	mainPath := path.Join(pathutils.GetModulePath("CrawlerDemo"), "main", "main.go")
	log.Println("work directory: ", mainPath)

	waiters := make([]chan int, n)
	outers := make([]bytes.Buffer, n)
	cmds := make([]*exec.Cmd, n)

	for i := 0; i < n; i++ {
		cmds[i] = exec.Command("go", "run", mainPath)
		cmds[i].Stdout = &outers[i]
		waiters[i] = make(chan int)

		log.Println("start process ", i+1)

		go execCmd(cmds[i], waiters[i])
	}

	// wait processes
	var needCheck bool = true

	waitChan := make(chan int)
	interruptChan := make(chan int)

	waitProc := func() {
		for _, waiter := range waiters {
			s := <-waiter
			if s == 0 {
				needCheck = false
			}
		}
		waitChan <- 0
	}
	go waitProc()

	// process interrupt
	go func() {
		err := listenRemoteInterrupt(interruptChan)
		if err != nil {
			log.Println(err)
		}
	}()

	select {
	case <-waitChan:
		{
			// check processes result
			if needCheck && checkOuters(outers) {
				fmt.Println("successed")
			} else {
				fmt.Println("failed")
			}
		}
	case <-interruptChan:
		{
			log.Println("receive interrupt")

			for _, cmd := range cmds {
				if cmd == nil {
					continue
				}
				if err := cmd.Process.Kill(); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func checkOuters(outers []bytes.Buffer) bool {
	n := len(outers)
	outs := make([][]string, n)

	for i := 0; i < n; i++ {
		outs[i] = strings.Split(outers[i].String(), "\n")
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

func listenRemoteInterrupt(interruptChan chan int) error {
	listen, err := net.Listen("tcp", ":2233")
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
