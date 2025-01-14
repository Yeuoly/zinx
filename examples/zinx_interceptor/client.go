package main

import (
	"github.com/Yeuoly/zinx/ziface"
	"github.com/Yeuoly/zinx/znet"
	"time"
)

func main() {
	// 客户端向服务端发送消息
	client := znet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(connection ziface.IConnection) {
		_ = connection.SendMsg(1, []byte("hello zinx"))
	})
	client.Start()
	time.Sleep(time.Second)
}
