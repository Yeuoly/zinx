package router

import (
	"github.com/Yeuoly/zinx/ziface"
	"github.com/Yeuoly/zinx/zlog"
	"github.com/Yeuoly/zinx/znet"
)

type HelloRouter struct {
	znet.BaseRouter
}

func (hr *HelloRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF(string(request.GetData()))
}
