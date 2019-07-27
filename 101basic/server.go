package main

import (
	"flag"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {

	flag.Parse()

	// 创建一个新的客户端对象
	s := server.NewServer()

	// 注册服务
	// 注册方法1: 可以采用 RegisterName方法来给服务命名一个自定义的名字, 比如 MyArith
	//s.RegisterName("MyArith", new(example.Arith), "")
	// 注册方法2: 使用 Arith (服务的默认名称为服务的Type名)
	s.Register(new(example.Arith), "")

	// 开启服务端, 开始对外提供服务, 等待客户端连接(底层实现的是 net.Listener.Accept(), 阻塞等待连接)
	s.Serve("tcp", *addr)
}
