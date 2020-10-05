package main

import (
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	"github.com/pantskun/golearn/customEtcdclt/etcdcltmicro/handler"
)

func main() {
	srv := service.New(
		service.Name("etcdservice"),
		service.Version("latest"),
	)

	srv.Handle(new(handler.EtcdcltMicro))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
