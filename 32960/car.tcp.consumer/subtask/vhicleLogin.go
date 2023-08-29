package subtask

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"car.tcp.consumer/conf"
	"car.tcp.consumer/entity/dbmodel"
	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/repo"

	"car.tcp.consumer/tools"
	"github.com/dapr/go-sdk/service/common"
	"github.com/wangcheng0509/gpkg/try"
)

// Topic注册
func VehicleLogin(s common.Service) {
	var vehicleLoginSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "VehicleLogin",
		Route:      "/VehicleLogin",
	}

	if err := s.AddTopicEventHandler(vehicleLoginSubscription, vehicleLoginHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

func VehicleLoginTest(req msgModel.Message) (retry bool, err error) {
	// 32960协议解析
	vlogin, vloginLog, err := vehicleLoginUnmarshal(req)
	if err != nil {
		return false, err
	}
	// 写入DB
	if err := repo.Vlogin(vlogin, vloginLog); err != nil {
		return true, err
	}
	return false, nil
}

// task
func vehicleLoginHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s", e.PubsubName, e.Topic, e.ID)
	// 解析请求msg
	var req msgModel.Message
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		fmt.Println("解析请求msg Error:", err.Error())
		return false, err
	}
	fmt.Println(req)
	// 32960协议解析
	v, vLog, err := vehicleLoginUnmarshal(req)
	log.Println("vehicleLogin:", v)
	log.Println("vehicleLoginLog:", vLog)
	if err != nil {
		fmt.Println("32960协议解析 Error:", err.Error())
		return false, err
	}
	// 写入DB
	if err := repo.Vlogin(v, vLog); err != nil {
		fmt.Println("写入DB Error:", err.Error())
		return true, err
	}
	return false, nil
}

func vehicleLoginUnmarshal(req msgModel.Message) (dbmodel.Vlogin, dbmodel.VloginLog, error) {
	var vlogin dbmodel.Vlogin
	var vloginLog dbmodel.VloginLog
	var err error
	if len(req.Data) < 30 {
		return vlogin, vloginLog, errors.New("param is error")
	}
	try.Try(func() {
		defer func() {
			if errInner := recover(); errInner != nil {
				try.Throw(1, errInner.(string))
			}
		}()
		vlogin.Vin = req.Vin
		vloginLog.Vin = req.Vin
		// 年月日时分秒
		if msgDate, err := tools.GetMsgDate(req.Data); err == nil {
			vlogin.Year = msgDate.Year
			vlogin.Month = msgDate.Month
			vlogin.Day = msgDate.Day
			vlogin.Hour = msgDate.Hour
			vlogin.Minutes = msgDate.Minutes
			vlogin.Seconds = msgDate.Seconds
			vloginLog.Year = int32(msgDate.Year)
			vloginLog.Month = int32(msgDate.Month)
			vloginLog.Day = int32(msgDate.Day)
			vloginLog.Hour = int32(msgDate.Hour)
			vloginLog.Minutes = int32(msgDate.Minutes)
			vloginLog.Seconds = int32(msgDate.Seconds)
		} else {
			try.Throw(1, "param is error")
		}
		// 登录流水号
		vlogin.Seq = tools.GetUInt16(req.Data[6:8])
		vloginLog.Seq = int32(tools.GetUInt16(req.Data[6:8]))
		// Iccid
		vlogin.IccId = string(req.Data[8 : 8+20])
		vloginLog.IccId = string(req.Data[8 : 8+20])
		// 可充电储能子系统数
		vlogin.Num = int(req.Data[28])
		vloginLog.Num = int32(req.Data[28])
		// 可充电储能系统编码长度
		vlogin.Length = int(req.Data[29])
		vloginLog.Length = int32(req.Data[29])
		// 可充电储能系统编码列表
		if len(req.Data) > 30 {
			vlogin.EnergyId = string(req.Data[30:len(req.Data)])
			vloginLog.EnergyId = string(req.Data[30:len(req.Data)])
		}
		vlogin.CreatedTime = time.Now()
		vloginLog.CreatedTime = time.Now()
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return vlogin, vloginLog, err
}
