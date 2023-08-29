package handle

import (
	"fmt"
	"time"

	"car.tcp.service/entity/enums"
	"car.tcp.service/handle/base"
	"car.tcp.service/handle/service"
	"car.tcp.service/tools"

	"github.com/aceld/zinx/znet"
	dapr "github.com/dapr/go-sdk/client"
)

func ServeStart() {
	var err error
	if err = MasterLogin(); err != nil {
		fmt.Println(err.Error())
		panic("MasterLogin Error")
	}
	base.DaprClient, err = dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer base.DaprClient.Close()
	routeRegister()
	base.TcpServer.Serve()
}

func routeRegister() {
	base.TcpServer = znet.NewServer()
	base.TcpServer.Use(tools.CheckCode)
	base.TcpServer.AddRouterSlices(1, service.VloginHandle)
	base.TcpServer.AddRouterSlices(2, service.InfoHandle)
	base.TcpServer.AddRouterSlices(3, service.ReinfoHandle)
	base.TcpServer.AddRouterSlices(4, service.VlogoutHandle)
	base.TcpServer.AddRouterSlices(enums.Msg.Heartbeat.Value, service.HeartbeatHandle)
	base.TcpServer.SetOnConnStop(base.DoConnectionLost)
	base.TcpServer.StartHeartBeat(100 * time.Second)
}

func MasterLogin() error {
	// TODO 注册TCP服务器
	return nil
}
