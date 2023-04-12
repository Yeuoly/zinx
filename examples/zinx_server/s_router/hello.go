package s_router

import (
	"github.com/Yeuoly/zinx/ziface"
	"github.com/Yeuoly/zinx/zlog"
	"github.com/Yeuoly/zinx/znet"
)

type HelloZinxRouter struct {
	znet.BaseRouter
}

// HelloZinxRouter Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	zlog.Debug("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	zlog.Debug("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(3, []byte("Hello Zinx Router[FromServer]"))
	if err != nil {
		zlog.Error(err)
	}
}
