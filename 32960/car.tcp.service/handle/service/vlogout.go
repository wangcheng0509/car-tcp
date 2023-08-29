package service

import (
	"fmt"

	"car.tcp.service/entity/enums"
	"car.tcp.service/entity/handleEntity"
	"car.tcp.service/handle/base"
	"car.tcp.service/repository"
	"github.com/aceld/zinx/ziface"
	"github.com/wangcheng0509/gpkg/gredis"
)

func VlogoutHandle(request ziface.IRequest) {
	msg := base.GetMessage(request.GetData())
	vlogout(request, msg)
}

func vlogout(request ziface.IRequest, msg handleEntity.Message) {
	fmt.Println("clent-", request.GetMsgID(), ",msg：", request.GetData())
	if err := vDelete(msg); err != nil {
		msg.CmdRsp = int(enums.Rsp.Error.Value)
		request.GetConnection().SendMsg(enums.Msg.VehicleLogout.Value, msg.Marshal())
		return
	}
	if err := base.DaprPub(enums.Msg.VehicleLogout.Name, msg); err != nil {
		msg.CmdRsp = int(enums.Rsp.Error.Value)
		request.GetConnection().SendMsg(enums.Msg.VehicleLogout.Value, msg.Marshal())
		return
	}

	msg.CmdRsp = int(enums.Rsp.Success.Value)
	request.GetConnection().SendMsg(enums.Msg.VehicleLogout.Value, msg.Marshal())
}

func vDelete(msg handleEntity.Message) error {
	// 链接管理器注册
	repository.OfflineByVin(msg.Vin)
	if err := gredis.Cluster.HDel(vloginRedisKey, msg.Vin).Err(); err != nil {
		return err
	}
	return nil
}
