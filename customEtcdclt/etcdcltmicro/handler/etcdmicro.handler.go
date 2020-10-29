package handler

import (
	"context"
	"path"

	"github.com/pantskun/commonutils/pathutils"
	proto "github.com/pantskun/golearn/customEtcdclt/etcdcltmicro/proto"
	"github.com/pantskun/golearn/customEtcdclt/etcdinteraction"
)

type EtcdcltMicro struct{}

var _ proto.EtcdcltMicroHandler = (*EtcdcltMicro)(nil)

func (*EtcdcltMicro) Call(ctx context.Context, req *proto.Request, rep *proto.Response) error {
	action := getEtcdAction(req)
	config := etcdinteraction.GetEtcdClientConfig(path.Join(pathutils.GetModulePath("customEtcdclt"), "configs/etcdClientConfig.json"))
	client := etcdinteraction.GetEtcdClient(config)
	msg, err := etcdinteraction.ExecuteAction(action, client)

	rep.Msg = msg
	if err != nil {
		rep.Err = err.Error()
	}

	return nil
}

func getEtcdAction(req *proto.Request) etcdinteraction.EtcdActionInterface {
	switch req.Action {
	case "get":
		return etcdinteraction.NewGetAction(req.Key)
	case "put":
		return etcdinteraction.NewPutAction(req.Key, req.Value)
	case "del":
		return etcdinteraction.NewDeleteAction(req.Key)
	}

	return nil
}
