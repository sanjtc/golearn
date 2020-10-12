package main

import (
	"log"
	"os"

	"github.com/pantskun/golearn/customEtcdclt/httpclient/client"
)

func main() {
	log.Println(client.HTTPClient(":8080", make(chan os.Signal, 1)))
}
