package main

import (
	"context"
	"fmt"

	"github.com/micro/micro/v3/service/client"
	pb "github.com/pantskun/golearn/customEtcdclt/etcdcltmicro/proto"
)

func main() {
	// req := client.NewRequest("etcdcltmicro", "EtcdcltMicro.Call", &pb.Request{Action: "get", Key: "key"})
	// rsp := new(pb.Response)

	// err := client.Call(context.Background(), req, &rsp)

	hw := pb.NewEtcdcltMicroService("etcdcltmicro", client.DefaultClient)
	rsp, err := hw.Call(context.Background(), &pb.Request{Action: "put", Key: "key", Value: "value"})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}
