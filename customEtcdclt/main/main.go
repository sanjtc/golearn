package main

import (
	"log"
	"os"

	"github.com/pantskun/golearn/customEtcdclt/cmdclient"
	"github.com/pantskun/golearn/customEtcdclt/httpclient"
)

func main() {
	const httpMode = true

	if httpMode {
		log.Println(httpclient.HTTPClient(":8080", make(chan os.Signal, 1)))
	} else {
		cmdclient.CMDClient()
	}
}
