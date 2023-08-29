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

func Vehicle(s common.Service) {
	var vehicleSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "Vehicle",
		Route:      "/Vehicle",
	}

	if err := s.AddTopicEventHandler(vehicleSubscription, vehicleHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

func VehicleTest(req msgModel.InfoCateModel) (retry bool, err error) {
	// 32960协议解析
	v, vLog, err := vehicleUnmarshal(req)
	if err != nil {
		return false, err
	}
	// 写入DB
	if err := repo.VehicleCreate(v, vLog); err != nil {
		return true, err
	}
	return false, nil
}

// task
func vehicleHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s", e.PubsubName, e.Topic, e.ID)
	// 解析请求msg
	var req msgModel.InfoCateModel
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		fmt.Println("解析请求msg Error:", err.Error())
		return false, err
	}
	fmt.Println(req)
	// 32960协议解析
	v, vLog, err := vehicleUnmarshal(req)
	log.Println("vehicle:", v)
	log.Println("vehicleLog:", vLog)
	if err != nil {
		fmt.Println("32960协议解析 Error:", err.Error())
		return false, err
	}
	// 写入DB
	if err := repo.VehicleCreate(v, vLog); err != nil {
		fmt.Println("写入DB Error:", err.Error())
		return true, err
	}
	return false, nil
}

func vehicleUnmarshal(req msgModel.InfoCateModel) (dbmodel.Vehicle, dbmodel.VehicleLog, error) {
	var v dbmodel.Vehicle
	var vLog dbmodel.VehicleLog
	var err error
	if len(req.Data) < 16 {
		return v, vLog, errors.New("param is error")
	}
	try.Try(func() {
		defer func() {
			if errInner := recover(); errInner != nil {
				try.Throw(1, errInner.(string))
			}
		}()
		v.Vin = req.Vin
		vLog.Vin = req.Vin
		// 年月日时分秒
		v.Year = req.MsgDate.Year
		v.Month = req.MsgDate.Month
		v.Day = req.MsgDate.Day
		v.Hour = req.MsgDate.Hour
		v.Minutes = req.MsgDate.Minutes
		v.Seconds = req.MsgDate.Seconds
		vLog.Year = int32(req.MsgDate.Year)
		vLog.Month = int32(req.MsgDate.Month)
		vLog.Day = int32(req.MsgDate.Day)
		vLog.Hour = int32(req.MsgDate.Hour)
		vLog.Minutes = int32(req.MsgDate.Minutes)
		vLog.Seconds = int32(req.MsgDate.Seconds)
		// 整车数据解析
		v.Status = int(req.Data[0])
		vLog.Status = int32(req.Data[0])
		v.Charging = int(req.Data[1])
		vLog.Charging = int32(req.Data[1])
		v.Mode = int(req.Data[2])
		vLog.Mode = int32(req.Data[2])
		v.Speed = tools.GetUInt16(req.Data[3:5])
		vLog.Speed = int32(tools.GetUInt16(req.Data[3:5]))
		v.Mileage = tools.GetUInt32(req.Data[5:9])
		vLog.Mileage = int32(tools.GetUInt32(req.Data[5:9]))
		v.Voltage = tools.GetUInt16(req.Data[9:11])
		vLog.Voltage = int32(tools.GetUInt16(req.Data[9:11]))
		v.Current = tools.GetUInt16(req.Data[11:13])
		vLog.Current = int32(tools.GetUInt16(req.Data[11:13]))
		v.Soc = int(req.Data[13])
		vLog.Soc = int32(req.Data[13])
		v.Dc = int(req.Data[14])
		vLog.Dc = int32(req.Data[14])
		v.Gear = int(req.Data[15])
		vLog.Gear = int32(req.Data[15])
		v.Resistance = tools.GetUInt16(req.Data[16:18])
		vLog.Resistance = int32(tools.GetUInt16(req.Data[16:18]))
		v.BrakePedal = int(req.Data[18])
		vLog.BrakePedal = int32(req.Data[18])
		v.AcceleratorPedal = int(req.Data[19])
		vLog.AcceleratorPedal = int32(req.Data[19])

		v.CreatedTime = time.Now()
		vLog.CreatedTime = time.Now()
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return v, vLog, err
}
