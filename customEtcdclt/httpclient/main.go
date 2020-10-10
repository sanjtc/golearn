package main

import (
	"log"
	"os"

	"github.com/pantskun/golearn/customEtcdclt/httpclient/httpclient"
)

func main() {
	log.Println(httpclient.HTTPClient(":8080", make(chan os.Signal, 1)))
}
