package main

import (
	"golearn/microServer/helloworld/handler"

	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
)

func main() {
	srv := service.New(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	srv.Handle(new(handler.Helloworld))
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
