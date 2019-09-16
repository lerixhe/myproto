package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"myproto/1.proto-test/prototext"
)

func main() {
	text := &prototext.TestRequest{
		Name:   "Lerix",
		Weight: []int32{120, 130, 124, 150, 190},
		Height: 180,
		Moto:   "面对问题，解决问题",
	}
	fmt.Println(text)
	data, err := proto.Marshal(text)
	if err != nil {
		fmt.Println("编码失败", err)
	}
	fmt.Println(data)
	newText := new(prototext.TestRequest)
	err = proto.Unmarshal(data, newText)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	fmt.Println(newText)
	fmt.Println(newText.Name)
	fmt.Println(newText.Height)
	fmt.Println(newText.Moto)
}
