package main

import (
	"context"
	"fmt"

	pd "myproto/3.grpc-test/prototext"

	"google.golang.org/grpc"
)

func main() {
	// 客户端创建连接，接入监听服务器
	conn, err := grpc.Dial("127.0.0.1:10086", grpc.WithInsecure())
	if err != nil {
		fmt.Println("网络异常", err)
	}
	// 退出前关闭连接
	defer conn.Close()
	// 从连接冲获取客户端实例
	client := pd.NewHelloserverClient(conn)
	// 客户端实例调用服务端的方法，并传入参数
	hellore, err := client.SayHello(context.Background(), &pd.HelloReq{Name: "熊猫"})
	if err != nil {
		fmt.Println("sayhello服务调用失败")
	}
	// 获得服务端的返回值
	fmt.Println("调用sayhello的返回", hellore.Msg)

	namere, err := client.SayName(context.Background(), &pd.NameReq{Name: "猩猩"})
	if err != nil {
		fmt.Println("sayname服务调用失败")
	}
	fmt.Println("调用sayname服务的返回", namere.Msg)
}
