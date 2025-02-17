package main

import (
	"fmt"
	"github.com/Yeuoly/zinx/examples/zinx_client/c_router"
	"github.com/Yeuoly/zinx/ziface"
	"github.com/Yeuoly/zinx/zlog"
	"github.com/Yeuoly/zinx/znet"
	"os"
	"os/signal"
	"time"
)

// 客户端自定义业务
func business(conn ziface.IConnection) {

	for {
		err := conn.SendMsg(100, []byte("Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
			zlog.Error(err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

// 创建连接的时候执行
func DoClientConnectedBegin(conn ziface.IConnection) {
	zlog.Debug("DoConnecionBegin is Called ... ")

	//设置两个链接属性，在连接创建之后
	conn.SetProperty("Name", "刘丹冰")
	conn.SetProperty("Home", "https://yuque.com/aceld")

	go business(conn)
}

// 连接断开的时候执行
func DoClientConnectedLost(conn ziface.IConnection) {
	//在连接销毁之前，查询conn的Name，Home属性
	if name, err := conn.GetProperty("Name"); err == nil {
		zlog.Debug("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		zlog.Debug("Conn Property Home = ", home)
	}

	zlog.Debug("DoClientConnectedLost is Called ... ")
}

func main() {
	//创建一个Client句柄，使用Zinx的API
	client := znet.NewClient("127.0.0.1", 9000)

	//添加首次建立链接时的业务
	client.SetOnConnStart(DoClientConnectedBegin)
	client.SetOnConnStop(DoClientConnectedLost)

	//注册收到服务器消息业务路由
	client.AddRouter(2, &c_router.PingRouter{})
	client.AddRouter(3, &c_router.HelloRouter{})

	//启动客户端client
	client.Start()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Println("===exit===", sig)
	// 清理客户端
	client.Stop()
	time.Sleep(time.Second * 2)
}
