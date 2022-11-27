package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// ping test custom route
type PingRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	//Read the data from the client first
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//To go back to write  "ping...ping...ping"
	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	//1 Create the server object
	s := znet.NewServer()

	//2 Configure user-defined routes and services
	s.AddRouter(0, &PingRouter{})

	//3 Start the service
	s.Serve()
}
