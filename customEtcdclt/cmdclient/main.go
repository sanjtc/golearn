package main

import (
	"fmt"

	"github.com/pantskun/golearn/customEtcdclt/cmdclient/client"
)

func main() {
	if msg, err := client.CMDClient(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}
