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

// DrivemotorTopic 驱动电机
func DrivemotorTopic(s common.Service) {
	var drivemotorSub = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "Drivemotor",
		Route:      "/Drivemotor",
	}

	if err := s.AddTopicEventHandler(drivemotorSub, drivemotorHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

// DrivemotorTest ..
func DrivemotorTest(ctx context.Context, req msgModel.InfoCateModel) (retry bool, err error) {
	var drivemotors model.Drivemotors
	var drivemotorsLog model.DrivemotorsLog
	drivemotors.Vin = req.Vin
	drivemotors.Date = ToModelDate(req.MsgDate)
	drivemotors.CreatedTime = time.Now()
	err = unmarshalDrivemotor(req.Data, &drivemotors)
	if err != nil {
		return false, err
	}
	drivemotorsLog.Drivemotors = drivemotors

	err = repo.DrivemotorUpsert(ctx, drivemotors)
	if err != nil {
		return false, err
	}
	return false, repo.DrivemotorLogCreate(ctx, drivemotorsLog)

}

func drivemotorHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s \n", e.PubsubName, e.Topic, e.ID, e.Data)
	// 解析请求msg
	var req msgModel.InfoCateModel
	if err := e.Struct(&req); err != nil {
		return false, err
	}
	log.Println(req)

	var drivemotors model.Drivemotors
	var drivemotorsLog model.DrivemotorsLog
	drivemotors.Vin = req.Vin
	drivemotors.Date = ToModelDate(req.MsgDate)
	drivemotors.CreatedTime = time.Now()
	err = unmarshalDrivemotor(req.Data, &drivemotors)
	if err != nil {
		return false, err
	}
	drivemotorsLog.Drivemotors = drivemotors

	err = repo.DrivemotorUpsert(ctx, drivemotors)
	if err != nil {
		return false, err
	}
	return false, repo.DrivemotorLogCreate(ctx, drivemotorsLog)
}

func unmarshalDrivemotor(data []byte, dms *model.Drivemotors) error {
	if len(data) < 13 {
		return fmt.Errorf("驱动电机数据格式不正确:%v", data)
	}
	dms.Number = int32(data[0])
	data = data[1:]
	if len(data)%12 != 0 {
		return fmt.Errorf("驱动电机数据格式不正确:%v", data)
	}

	var drivemotors []*model.Drivemotor
	for i := 0; i < len(data)/12; i++ {
		var drivemotor model.Drivemotor
		drivemotor.No = int32(data[0])
		drivemotor.Status = int32(data[1])
		drivemotor.CtrlTemp = int32(data[2])
		drivemotor.Rotating = int32(binary.BigEndian.Uint16(data[3:5]))
		drivemotor.Torque = int32(binary.BigEndian.Uint16(data[5:7]))
		drivemotor.MotorTemp = int32(data[7])
		drivemotor.InputVoltage = int32(binary.BigEndian.Uint16(data[8:10]))
		drivemotor.DcBusCurrent = int32(binary.BigEndian.Uint16(data[10:12]))
		drivemotors = append(drivemotors, &drivemotor)
	}
	encoded, err := json.Marshal(&drivemotors)
	if err != nil {
		return err
	}
	dms.Data = string(encoded)

	return nil
}
