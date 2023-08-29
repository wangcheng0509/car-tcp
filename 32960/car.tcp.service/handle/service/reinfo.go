package service

import (
	"fmt"

	"car.tcp.service/entity/enums"
	"car.tcp.service/entity/handleEntity"
	"car.tcp.service/handle/base"
	"github.com/aceld/zinx/ziface"
)

func ReinfoHandle(request ziface.IRequest) {
	msg := base.GetMessage(request.GetData())
	info(request, msg)
}

func reinfo(request ziface.IRequest, msg handleEntity.Message) {
	fmt.Println("clent-", request.GetMsgID(), ",msg：", request.GetData())
	if err := base.DaprPub(enums.Msg.ResendInfo.Name, msg); err != nil {
		msg.CmdRsp = int(enums.Rsp.Error.Value)
		request.GetConnection().SendMsg(enums.Msg.ResendInfo.Value, msg.Marshal())
		return
	}
	msg.CmdRsp = int(enums.Rsp.Success.Value)
	request.GetConnection().SendMsg(enums.Msg.ResendInfo.Value, msg.Marshal())
}
