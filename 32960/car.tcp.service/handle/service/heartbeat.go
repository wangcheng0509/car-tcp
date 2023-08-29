package service

import (
	"fmt"

	"car.tcp.service/entity/enums"
	"car.tcp.service/handle/base"
	"github.com/aceld/zinx/ziface"
)

func HeartbeatHandle(request ziface.IRequest) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	msg := base.GetMessage(request.GetData())
	if err := vRegister(msg, request.GetConnection().GetConnID()); err != nil {
		msg.CmdRsp = int(enums.Rsp.Error.Value)
	} else {
		msg.CmdRsp = int(enums.Rsp.Success.Value)
	}
	if err := base.DaprPub(enums.Msg.Heartbeat.Name, msg); err != nil {
		fmt.Println(err.Error())
		msg.CmdRsp = int(enums.Rsp.Error.Value)
	}
	if err := request.GetConnection().SendMsg(enums.Msg.Heartbeat.Value, msg.Marshal()); err != nil {
		fmt.Println(err.Error())
	}
}
