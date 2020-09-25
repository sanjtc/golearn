package main

import (
	"log"

	"github.com/pantskun/golearn/customEtcdclt/cmdclient"
	"github.com/pantskun/golearn/customEtcdclt/httpclient"
)

func main() {
	const httpMode = true

	if httpMode {
		log.Println(httpclient.HTTPClient(":8080"))
	} else {
		cmdclient.CMDClient()
	}
}
