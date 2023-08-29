package subtask

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"car.tcp.consumer/conf"
	"car.tcp.consumer/entity/model"
	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/repo"
	"github.com/dapr/go-sdk/service/common"
)

// ChargeableVoltage 可充电储能装置电压
type ChargeableVoltage struct {
	SubSysNo               int   `json:"subSysNo" gorm:"column:subSysNo"`                             // 可充电储能子系统号 有效值范围 1~250
	ChargeableVoltage      int   `json:"chargeableVoltage" gorm:"column:chargeableVoltage"`           // 可充电储能装置电压 有效值范围围:0~10,000(表示0V~10000V)，最小计量单元:0.1V "0xFE0xFE"表示异常， ，"OxFFOxFF”表示无效
	ChargeableCurrent      int   `json:"chargeableCurrent" gorm:"column:chargeableCurrent"`           // 可充电储能装置电流 有效值范围: 0-20,000(数值偏移量1000A，表示 -1000A~+1000A)，最小计量单0.1A "0xFE0xFE"表示异常， ，"OxFFOxFF”表示无效
	SingleBatteryCount     int   `json:"singleBatteryCount" gorm:"column:singleBatteryCount"`         // N个电池单体 ，范围: 1~65531 "0xFE0xFE"表示异常， ，"OxFFOxFF”表示无效
	StartFrameBatteryNo    int   `json:"startFrameBatteryNo" gorm:"column:startFrameBatteryNo"`       // 当本帧单体个数别超过200时，应拆分成多帧m数据进行传输，有效值范围 1~65531
	TotalFrameBatteryCount int   `json:"totalFrameBatteryCount" gorm:"column:totalFrameBatteryCount"` // 本帧单体总数m;有效值范围:1~200
	SingleBatteryVotage    []int `json:"singleBatteryVotage" gorm:"column:singleBatteryVotage"`       // 有效范围:0-60,000 (表示0V-60.000V)  最小计量单元0.001V 单体电池电压个数等于本帧单体电池总数m，
}

// ChargeableVoltageTopic 发动机
func ChargeableVoltageTopic(s common.Service) {
	var chargeableVoltageSub = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "ChargeableVoltage",
		Route:      "/ChargeableVoltage",
	}

	if err := s.AddTopicEventHandler(chargeableVoltageSub, chargeableVoltageHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

// ChargeableVoltageTest ..
func ChargeableVoltageTest(ctx context.Context, req msgModel.InfoCateModel) (retry bool, err error) {
	var chargeableVoltage model.ChargeableVoltage
	var chargeableVoltageLog model.ChargeableVoltageLog
	chargeableVoltage.Vin = req.Vin
	chargeableVoltage.Date = ToModelDate(req.MsgDate)
	chargeableVoltage.CreatedTime = time.Now()
	err = unmarshalChargeableVoltage(req.Data, &chargeableVoltage)
	if err != nil {
		return false, err
	}
	chargeableVoltageLog.ChargeableVoltage = chargeableVoltage

	err = repo.ChargeableVoltageUpsert(ctx, chargeableVoltage)
	if err != nil {
		return false, err
	}
	return false, repo.ChargeableVoltageLogCreate(ctx, chargeableVoltageLog)

}

func chargeableVoltageHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s \n", e.PubsubName, e.Topic, e.ID, e.Data)
	// 解析请求msg
	var req msgModel.InfoCateModel
	if err := e.Struct(&req); err != nil {
		return false, err
	}
	log.Printf("consume: %v\n", req)

	var chargeableVoltage model.ChargeableVoltage
	var chargeableVoltageLog model.ChargeableVoltageLog
	chargeableVoltage.Vin = req.Vin
	chargeableVoltage.Date = ToModelDate(req.MsgDate)
	chargeableVoltage.CreatedTime = time.Now()
	err = unmarshalChargeableVoltage(req.Data, &chargeableVoltage)
	if err != nil {
		return false, err
	}
	chargeableVoltageLog.ChargeableVoltage = chargeableVoltage

	err = repo.ChargeableVoltageUpsert(ctx, chargeableVoltage)
	if err != nil {
		return false, err
	}
	return false, repo.ChargeableVoltageLogCreate(ctx, chargeableVoltageLog)
}

func unmarshalChargeableVoltage(data []byte, chargeableVoltage *model.ChargeableVoltage) error {
	if len(data) < 12 || len(data) != int(1+data[0]*(10+data[10]*2)) {
		return fmt.Errorf("充电储能子系统电压数据:%v", data)
	}

	var cvs []*ChargeableVoltage
	chargeableVoltage.Number = int32(data[0])
	data = data[1:]
	for i := 0; i < int(chargeableVoltage.Number); i++ {
		var cv ChargeableVoltage
		cv.SubSysNo = int(data[0])
		cv.ChargeableVoltage = int(binary.BigEndian.Uint16(data[1:3]))
		cv.ChargeableCurrent = int(binary.BigEndian.Uint16(data[3:5]))
		cv.SingleBatteryCount = int(binary.BigEndian.Uint16(data[5:7]))
		cv.StartFrameBatteryNo = int(binary.BigEndian.Uint16(data[7:9]))
		cv.TotalFrameBatteryCount = int(data[9])

		data = data[10:]
		for i := 0; i < cv.TotalFrameBatteryCount; i++ {
			cv.SingleBatteryVotage = append(cv.SingleBatteryVotage, int(binary.BigEndian.Uint16(data[0:2])))
			data = data[2:]
		}
		cvs = append(cvs, &cv)
	}

	encoded, err := json.Marshal(&cvs)
	if err != nil {
		return err
	}
	chargeableVoltage.Data = string(encoded)
	return nil
}
