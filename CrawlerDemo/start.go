package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"

	"github.com/pantskun/pathlib"
)

const n = 6

func main() {
	// start etcd
	etcdCmd := exec.Command("etcd")

	go func() {
		err := etcdCmd.Run()
		if err != nil {
			log.Println(err)
		}
	}()

	// start n processes
	execCmd := func(cmd *exec.Cmd, s chan int) {
		err := cmd.Run()
		if err != nil {
			log.Println(err)
			s <- 0
		} else {
			s <- 1
		}
	}

	mainPath := pathlib.GetModulePath("CrawlerDemo") + "/main/main.go"
	waiters := make([]chan int, n)
	outers := make([]bytes.Buffer, n)

	for i := 0; i < n; i++ {
		cmd := exec.Command("go", "run", mainPath)
		cmd.Stdout = &outers[i]
		waiters[i] = make(chan int)

		go execCmd(cmd, waiters[i])
	}

	// wait processes
	var needCheck bool = true

	for _, waiter := range waiters {
		s := <-waiter
		if s == 0 {
			needCheck = false
		}
	}

	// clean url data
	etcdctlCmd := exec.Command("etcdctl", "del", "--prefix", "https://")

	var etcdctlOuter bytes.Buffer
	etcdctlCmd.Stdout = &etcdctlOuter

	err := etcdctlCmd.Run()
	if err != nil {
		log.Println(err)
	}

	// close etcd
	_ = etcdCmd.Process.Kill()

	// check processes result
	if needCheck && checkOuters(outers) {
		log.Println("successed")
	} else {
		log.Println("failed")
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
				log.Println(url)
				return false
			}
		}
	}

	return true
}
