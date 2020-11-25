package main

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/pantskun/commonutils/osutils"
	"github.com/stretchr/testify/assert"
)

func TestProcessRemoteInterrupt(t *testing.T) {
	interruptChan := make(chan int)
	processRemoteInterrupt(":2233", interruptChan)

	if err := sendInterruptMsg(":2233"); err != nil {
		t.Log(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		t.Fatal("interrupt timeout")
	case <-interruptChan:
		return
	}
}

func sendInterruptMsg(addr string) error {
	conn, err := net.Dial("tcp", addr)
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

func TestWaitingResult(t *testing.T) {
	type TestCase struct {
		waitChen      chan int
		interruptChan chan int
		cmds          []osutils.Command
		expected      string
	}

	emptyChan := make(chan int, 1)
	noemptyChan := make(chan int, 1)

	cmd1 := osutils.NewCommand("echo", "test1")
	cmd2 := osutils.NewCommand("echo", "test2")
	cmd3 := osutils.NewCommand("echo", "test1")

	cmd1.Run()
	cmd2.Run()
	cmd3.Run()

	trueCmds := []osutils.Command{cmd1, cmd2}
	falseCmds := []osutils.Command{cmd1, cmd3}

	testCases := []TestCase{
		{waitChen: noemptyChan, interruptChan: emptyChan, cmds: trueCmds, expected: "successed"},
		{waitChen: noemptyChan, interruptChan: emptyChan, cmds: falseCmds, expected: "failed"},
		{waitChen: emptyChan, interruptChan: noemptyChan, cmds: falseCmds, expected: "interrupt"},
	}

	for _, testCase := range testCases {
		noemptyChan <- 1

		got := waitingResult(testCase.waitChen, testCase.interruptChan, testCase.cmds)
		assert.Equal(t, testCase.expected, got)
	}

}
