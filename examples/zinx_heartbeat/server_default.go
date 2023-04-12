package main

import (
	"github.com/Yeuoly/zinx/znet"
	"time"
)

func main() {
	s := znet.NewServer()

	//启动心跳检测
	s.StartHeartBeat(5 * time.Second)

	s.Serve()
}
