package main

import (
	"context"
	"fmt"
	"net"

	pd "myproto/3.grpc-test/prototext"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// 创建网络监听服务
	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误", err)
	}
	// 创建grpc服务
	srv := grpc.NewServer()
	// 注册grpc服务，并将server对象传入注册好的服务。
	pd.RegisterHelloserverServer(srv, &server{})
	// 运行整个服务
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络错误", err)
	}
}

// 实现HelloserverServer接口的两个方法
func (s *server) SayHello(context.Context, *pd.HelloReq) (*pd.HelloRsp, error) {
	return &pd.HelloRsp{Msg: "get Hello"}, nil
}

// 一个说名字的函数
func (s *server) SayName(context.Context, *pd.NameReq) (*pd.NameRsp, error) {
	return &pd.NameRsp{Msg: "get name"}, nil
}
