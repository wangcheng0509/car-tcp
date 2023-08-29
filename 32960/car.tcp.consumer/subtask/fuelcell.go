package subtask

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"car.tcp.consumer/conf"
	"car.tcp.consumer/entity/dbmodel"
	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/repo"

	"car.tcp.consumer/tools"
	"github.com/dapr/go-sdk/service/common"
	"github.com/wangcheng0509/gpkg/try"
)

func Fuelcell(s common.Service) {
	var fuelcellSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "Fuelcell",
		Route:      "/Fuelcell",
	}

	if err := s.AddTopicEventHandler(fuelcellSubscription, fuelcellHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

func FuelcellTest(req msgModel.InfoCateModel) (retry bool, err error) {
	// 32960协议解析
	v, vLog, err := fuelcellUnmarshal(req)
	if err != nil {
		return false, err
	}
	// 写入DB
	if err := repo.FuelcellCreate(v, vLog); err != nil {
		return true, err
	}
	return false, nil
}

// task
func fuelcellHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s", e.PubsubName, e.Topic, e.ID)
	// 解析请求msg
	var req msgModel.InfoCateModel
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		fmt.Println("解析请求msg Error:", err.Error())
		return false, err
	}
	fmt.Println(req)
	// 32960协议解析
	v, vLog, err := fuelcellUnmarshal(req)
	log.Println("fuelcell:", v)
	log.Println("fuelcellLog:", vLog)
	if err != nil {
		fmt.Println("32960协议解析 Error:", err.Error())
		return false, err
	}
	// 写入DB
	if err := repo.FuelcellCreate(v, vLog); err != nil {
		fmt.Println("写入DB Error:", err.Error())
		return true, err
	}
	return false, nil
}

func fuelcellUnmarshal(req msgModel.InfoCateModel) (dbmodel.Fuelcell, dbmodel.FuelcellLog, error) {
	var v dbmodel.Fuelcell
	var vLog dbmodel.FuelcellLog
	var err error
	if len(req.Data) < 18 {
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
		v.CellVoltage = tools.GetUInt16(req.Data[0:2])
		vLog.CellVoltage = int32(tools.GetUInt16(req.Data[0:2]))
		v.CellCurrent = tools.GetUInt16(req.Data[2:4])
		vLog.CellCurrent = int32(tools.GetUInt16(req.Data[2:4]))
		v.FuelConsumption = tools.GetUInt16(req.Data[4:6])
		vLog.FuelConsumption = int32(tools.GetUInt16(req.Data[4:6]))
		v.ProbeNum = tools.GetUInt16(req.Data[6:8])
		vLog.ProbeNum = int32(tools.GetUInt16(req.Data[6:8]))
		var probeTemps string
		for i := 8; i < 8+v.ProbeNum; i++ {
			probeTemps += strconv.Itoa(int(req.Data[i])) + ","
		}
		v.ProbeTemps = strings.Trim(probeTemps, ",")
		vLog.ProbeTemps = strings.Trim(probeTemps, ",")
		v.H_MaxTemp = tools.GetUInt16(req.Data[8+v.ProbeNum : 10+v.ProbeNum])
		vLog.H_MaxTemp = int32(tools.GetUInt16(req.Data[8+v.ProbeNum : 10+v.ProbeNum]))
		v.H_TempProbeCode = int(req.Data[10+v.ProbeNum])
		vLog.H_TempProbeCode = int32(req.Data[10+v.ProbeNum])
		v.H_MaxConc = tools.GetUInt16(req.Data[11+v.ProbeNum : 13+v.ProbeNum])
		vLog.H_MaxConc = int32(tools.GetUInt16(req.Data[11+v.ProbeNum : 13+v.ProbeNum]))
		v.H_ConcSensorCode = int(req.Data[13+v.ProbeNum])
		vLog.H_ConcSensorCode = int32(req.Data[13+v.ProbeNum])
		v.H_MaxPress = tools.GetUInt16(req.Data[14+v.ProbeNum : 16+v.ProbeNum])
		vLog.H_MaxPress = int32(tools.GetUInt16(req.Data[14+v.ProbeNum : 16+v.ProbeNum]))
		v.H_PressSensorCode = int(req.Data[16+v.ProbeNum])
		vLog.H_PressSensorCode = int32(req.Data[16+v.ProbeNum])
		v.DcStatus = int(req.Data[17+v.ProbeNum])
		vLog.DcStatus = int32(req.Data[17+v.ProbeNum])

		v.CreatedTime = time.Now()
		vLog.CreatedTime = time.Now()
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return v, vLog, err
}
