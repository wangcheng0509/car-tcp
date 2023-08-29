package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"car.tcp.service/conf"
	"car.tcp.service/entity/enums"
	"car.tcp.service/entity/handleEntity"
	"car.tcp.service/handle/base"
	"car.tcp.service/repository"

	"github.com/aceld/zinx/ziface"
	"github.com/wangcheng0509/gpkg/gredis"
)

const vloginRedisKey = "car-client-conn"

func VloginHandle(request ziface.IRequest) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	msg := base.GetMessage(request.GetData())
	if err := vlogin(request, msg); err != nil {
		msg.CmdRsp = int(enums.Rsp.Error.Value)
	} else {
		msg.CmdRsp = int(enums.Rsp.Success.Value)
	}
	if err := request.GetConnection().SendMsg(enums.Msg.VehicleLogin.Value, msg.Marshal()); err != nil {
		fmt.Println(err.Error())
	}
}

func vlogin(request ziface.IRequest, msg handleEntity.Message) error {
	fmt.Println("clent-", request.GetMsgID(), ",msg：", request.GetData(), ";", string(request.GetData()))
	// 客户端链接注册
	err := vRegister(msg, request.GetConnection().GetConnID())
	if err != nil {
		return err
	}
	// 数据推送
	b, _ := json.Marshal(msg)
	fmt.Println(b)
	fmt.Println(string(b))
	if err := base.DaprPub(enums.Msg.VehicleLogin.Name, msg); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func vRegister(msg handleEntity.Message, connStr uint64) error {
	fmt.Println("vRegister begin")
	// 链接管理器注册
	cache := handleEntity.CacheModel{
		Vin:     msg.Vin,
		Host:    conf.Conf.Local.LocalHost,
		Port:    "8080",
		ConnStr: int(connStr),
	}
	fmt.Println("vRegister DB begin")
	if err := repository.CreateConn(cache); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("vRegister Cache begin")
	cacheByte, _ := cache.MarshalBinary()
	if err := gredis.Client.HSet(vloginRedisKey, msg.Vin, cacheByte).Err(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func getMsg(msg handleEntity.Message) (handleEntity.VloginModel, error) {
	reqMsg := handleEntity.VloginModel{}
	var err error
	defer func() {
		if err := recover(); err != nil {
			reqMsg = handleEntity.VloginModel{}
		}
	}()
	if len(msg.Data) < 30 {
		err = errors.New("非法入参")
		return reqMsg, err
	}
	reqMsg.Year = msg.Data[0]
	reqMsg.Month = msg.Data[1]
	reqMsg.Day = msg.Data[2]
	reqMsg.Hour = msg.Data[3]
	reqMsg.Minutes = msg.Data[4]
	reqMsg.Seconds = msg.Data[5]
	// TODO 计算登入流水号
	// reqMsg.LoginNum =
	reqMsg.IccId = string(msg.Data[8 : 8+21])
	reqMsg.SubSystemNumber = msg.Data[29]
	reqMsg.SystemEncodeLength = msg.Data[30]
	reqMsg.SystemEncode = string(msg.Data[31:])
	return reqMsg, err
}
