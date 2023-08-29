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

func Extreme(s common.Service) {
	var extremeSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "Extreme",
		Route:      "/Extreme",
	}

	if err := s.AddTopicEventHandler(extremeSubscription, extremeHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

func ExtremeTest(req msgModel.InfoCateModel) (retry bool, err error) {
	// 32960协议解析
	v, vLog, err := extremeUnmarshal(req)
	if err != nil {
		return false, err
	}
	// 写入DB
	if err := repo.ExtremeCreate(v, vLog); err != nil {
		return true, err
	}
	return false, nil
}

// task
func extremeHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s", e.PubsubName, e.Topic, e.ID)
	log.Println("")
	// 解析请求msg
	var req msgModel.InfoCateModel
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		return false, err
	}
	fmt.Println(req)
	// 32960协议解析
	v, vLog, err := extremeUnmarshal(req)
	log.Println("extreme:", v)
	log.Println("extremeLog:", vLog)
	if err != nil {
		fmt.Println("32960协议解析 Error:", err.Error())
		return false, err
	}
	// 写入DB
	if err := repo.ExtremeCreate(v, vLog); err != nil {
		fmt.Println("写入DB Error:", err.Error())
		return true, err
	}
	return false, nil
}

func extremeUnmarshal(req msgModel.InfoCateModel) (dbmodel.Extreme, dbmodel.ExtremeLog, error) {
	var v dbmodel.Extreme
	var vLog dbmodel.ExtremeLog
	var err error
	if len(req.Data) < 14 {
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
		v.MaxVoltageBatterySubsysNo = int(req.Data[0])
		vLog.MaxVoltageBatterySubsysNo = int32(req.Data[0])
		v.MaxVoltageBatteryCode = int(req.Data[1])
		vLog.MaxVoltageBatteryCode = int32(req.Data[1])
		v.MaxBatteryVoltage = tools.GetUInt16(req.Data[2:4])
		vLog.MaxBatteryVoltage = int32(tools.GetUInt16(req.Data[2:4]))
		v.MinVoltageBatterySubsysNo = int(req.Data[4])
		vLog.MinVoltageBatterySubsysNo = int32(req.Data[4])
		v.MinVoltageBatteryCode = int(req.Data[5])
		vLog.MinVoltageBatteryCode = int32(req.Data[5])
		v.MinBatteryVoltage = tools.GetUInt16(req.Data[6:8])
		vLog.MinBatteryVoltage = int32(tools.GetUInt16(req.Data[6:8]))
		v.MaxTempSubsysNo = int(req.Data[8])
		vLog.MaxTempSubsysNo = int32(req.Data[8])
		v.MaxTempProbeNo = int(req.Data[9])
		vLog.MaxTempProbeNo = int32(req.Data[9])
		v.MaxTemp = int(req.Data[10])
		vLog.MaxTemp = int32(req.Data[10])
		v.MinTempSubsysNo = int(req.Data[11])
		vLog.MinTempSubsysNo = int32(req.Data[11])
		v.MinTempProbeNo = int(req.Data[12])
		vLog.MinTempProbeNo = int32(req.Data[12])
		v.MinTemp = int(req.Data[13])
		vLog.MinTemp = int32(req.Data[13])

		v.CreatedTime = time.Now()
		vLog.CreatedTime = time.Now()
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return v, vLog, err
}
