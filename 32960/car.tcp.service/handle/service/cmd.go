package service

import (
	"car.tcp.service/entity/enums"
	"car.tcp.service/entity/handleEntity"
	"car.tcp.service/handle/base"
)

func CmdDown(req handleEntity.CmdModel) (int, error) {
	if conn, err := base.TcpServer.GetConnMgr().Get(req.ConnId); err != nil {
		return 1, err
	} else {
		if err := conn.SendMsg(enums.Msg.Cmd.Value, req.Msg); err != nil {
			return 2, err
		}
	}
	return 0, nil
}
