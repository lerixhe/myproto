package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("网络链接失败1", err)
	}
	var pd int
	// Call func(serviceMethod string, args interface{}, reply interface{}) error
	err = cli.Call("Lerix.GetInfo", 2, &pd)
	if err != nil {
		fmt.Println("网络链接失败2", err)
	}
	fmt.Println("服务端发来的值：", pd)
}
