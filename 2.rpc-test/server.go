package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

// func lerixtext(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "hello world,lerix")
// }

/*
- 方法是导出的
- 方法有两个参数，都是导出类型或内建类型
- 方法的第二个参数是指针
- 方法只有一个error接口类型的返回值
事实上，方法必须看起来像这样：

func (t *T) MethodName(argType T1, replyType *T2) error
*/

type Lerix int

func (t *Lerix) GetInfo(argType int, replyType *int) error {
	fmt.Println("客户端发来信息：", argType)
	*replyType = argType + 1
	return nil
}
func main() {

	// http.HandleFunc("/lerix", lerixtext)
	pd := new(Lerix)
	rpc.Register(pd)
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误")
	}
	http.Serve(ln, nil)
	// http.Handle()
}
