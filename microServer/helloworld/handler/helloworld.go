package handler

import (
	"context"

	"github.com/micro/go-micro/v3/logger"

	helloworld "github.com/pantskun/golearn/microServer/helloworld/proto"
)

type Helloworld struct{}

func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	logger.Info("Received Helloworld.Call request")

	rsp.Msg = "Hello " + req.Name

	return nil
}
