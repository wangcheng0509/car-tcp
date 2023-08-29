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

func Location(s common.Service) {
	var vehicleSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "Location",
		Route:      "/Location",
	}

	if err := s.AddTopicEventHandler(vehicleSubscription, locationHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

// task
func locationHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s", e.PubsubName, e.Topic, e.ID)
	// 解析请求msg
	var req msgModel.InfoCateModel
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		fmt.Println("解析请求msg Error:", err.Error())
		return false, err
	}
	fmt.Println(req)
	// 32960协议解析
	v, vLog, err := locationUnmarshal(req)
	log.Println("location:", v)
	log.Println("location:", vLog)
	if err != nil {
		fmt.Println("32960协议解析 Error:", err.Error())
		return false, err
	}
	// 写入DB
	if err := repo.LocationCreate(v, vLog); err != nil {
		fmt.Println("写入DB Error:", err.Error())
		return true, err
	}
	return false, nil
}

func locationUnmarshal(req msgModel.InfoCateModel) (dbmodel.Location, dbmodel.LocationLog, error) {
	var v dbmodel.Location
	var vLog dbmodel.LocationLog
	var err error
	if len(req.Data) < 9 {
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
		// 位置数据解析
		v.Status = int(req.Data[0])
		vLog.Status = int32(req.Data[0])
		v.Longitude = tools.GetUInt32(req.Data[1:5])
		vLog.Longitude = int32(tools.GetUInt32(req.Data[1:5]))
		v.Latitude = tools.GetUInt32(req.Data[5:9])
		vLog.Latitude = int32(tools.GetUInt32(req.Data[5:9]))

		v.CreatedTime = time.Now()
		vLog.CreatedTime = time.Now()
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return v, vLog, err
}
