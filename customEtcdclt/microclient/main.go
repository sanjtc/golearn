package main

import (
	"context"
	"fmt"

	"github.com/micro/micro/v3/service/client"
	pb "github.com/pantskun/golearn/customEtcdclt/etcdcltmicro/proto"
)

func main() {
	ms := pb.NewEtcdcltMicroService("etcdcltmicro", client.DefaultClient)

	rsp, err := ms.Call(context.Background(), &pb.Request{Action: "put", Key: "key", Value: "value"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}
