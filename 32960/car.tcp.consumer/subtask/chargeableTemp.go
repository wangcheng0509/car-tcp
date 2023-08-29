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

// ChargeableTemp 充电储能装置温度数据
type ChargeableTemp struct {
	SubSysNo       int   `json:"subSysNo" gorm:"column:subSysNo"`             // 可充电储能子系统号
	ProbeCount     int   `json:"probeCount" gorm:"column:probeCount"`         // 可充电储能 温度探针个数 N个温度探针，有效值范围:1~65531,"OxFF,OxFE”表示异常，"OxFF,OxFF”表示无效
	ProbeTempValue []int `json:"probeTempValue" gorm:"column:probeTempValue"` // 可充电储能子系统 各温度探针检 测到的温度值 范围:0~250 (数据偏移量40'C 表示-40～+210'C)，最小计量单元:1°C,"OxFE”表示异常，"OxFF”表示无效
}

// ChargeableTempTopic 个可充电储能子系统上温度
func ChargeableTempTopic(s common.Service) {
	var chargeableTempSub = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "ChargeableTemp",
		Route:      "/ChargeableTemp",
	}

	if err := s.AddTopicEventHandler(chargeableTempSub, chargeableTempHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

// ChargeableTempTest ..
func ChargeableTempTest(ctx context.Context, req msgModel.InfoCateModel) (retry bool, err error) {
	var chargeableTemp model.ChargeableTemp
	var chargeableTempLog model.ChargeableTempLog
	chargeableTemp.Vin = req.Vin
	chargeableTemp.Date = ToModelDate(req.MsgDate)
	err = unmarshalChargeableTemp(req.Data, &chargeableTemp)
	if err != nil {
		return false, err
	}
	chargeableTempLog.ChargeableTemp = chargeableTemp

	err = repo.ChargeableTempUpsert(ctx, chargeableTemp)
	if err != nil {
		return false, err
	}
	return false, repo.ChargeableTempLogCreate(ctx, chargeableTempLog)

}

func chargeableTempHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s \n", e.PubsubName, e.Topic, e.ID, e.Data)
	// 解析请求msg
	var req msgModel.InfoCateModel
	if err := e.Struct(&req); err != nil {
		return false, err
	}
	log.Printf("consume: %v\n", req)

	var chargeableTemp model.ChargeableTemp
	var chargeableTempLog model.ChargeableTempLog
	chargeableTemp.Vin = req.Vin
	chargeableTemp.Date = ToModelDate(req.MsgDate)
	err = unmarshalChargeableTemp(req.Data, &chargeableTemp)
	if err != nil {
		return false, err
	}
	chargeableTempLog.ChargeableTemp = chargeableTemp

	err = repo.ChargeableTempUpsert(ctx, chargeableTemp)
	if err != nil {
		return false, err
	}
	return false, repo.ChargeableTempLogCreate(ctx, chargeableTempLog)

}

func unmarshalChargeableTemp(data []byte, v *model.ChargeableTemp) error {
	if len(data) < 4 || len(data) != 1+int(data[0])*int(3+binary.BigEndian.Uint16(data[2:4])) {
		return fmt.Errorf("充电储能子系统温度数据:%v", data)
	}
	var cts []*ChargeableTemp
	v.Number = int32(data[0])
	v.CreatedTime = time.Now()

	data = data[1:]
	for i := 0; i < int(v.Number); i++ {
		var ct ChargeableTemp
		ct.SubSysNo = int(data[0])
		ct.ProbeCount = int(binary.BigEndian.Uint16(data[1:3]))
		data = data[3:]
		for i := 0; i < ct.ProbeCount; i++ {
			ct.ProbeTempValue = append(ct.ProbeTempValue, int(data[0]))
			data = data[1:]
		}
		cts = append(cts, &ct)
	}

	encoded, err := json.Marshal(&cts)
	if err != nil {
		return err
	}
	v.Data = string(encoded)
	return nil
}
