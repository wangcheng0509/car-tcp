package subtask

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"car.tcp.consumer/conf"
	"car.tcp.consumer/entity/model"
	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/repo"
	"github.com/dapr/go-sdk/service/common"
)

// EngineTopic 发动机
func EngineTopic(s common.Service) {
	var engineSub = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "Engine",
		Route:      "/Engine",
	}

	if err := s.AddTopicEventHandler(engineSub, engineHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

// ToModelDate .
func ToModelDate(date msgModel.MsgDate) model.Date {
	return model.Date{
		Year:    int32(date.Year),
		Month:   int32(date.Month),
		Day:     int32(date.Day),
		Hour:    int32(date.Hour),
		Minutes: int32(date.Minutes),
		Seconds: int32(date.Seconds),
	}
}

// EngineTest ..
func EngineTest(ctx context.Context, req msgModel.InfoCateModel) (retry bool, err error) {
	var engine model.Engine
	var engineLog model.EngineLog
	engine.Date = ToModelDate(req.MsgDate)
	engine.Vin = req.Vin
	engine.CreatedTime = time.Now()
	err = unmarshalEngine(req.Data, &engine)
	if err != nil {
		return false, err
	}
	engineLog.Engine = engine

	err = repo.EngineUpsert(ctx, engine)
	if err != nil {
		return false, err
	}
	return false, repo.EngineLogCreate(ctx, engineLog)
}

func engineHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s \n", e.PubsubName, e.Topic, e.ID, e.Data)
	// 解析请求msg
	var req msgModel.InfoCateModel
	if err := e.Struct(&req); err != nil {
		return false, err
	}
	log.Println(req)

	var engine model.Engine
	var engineLog model.EngineLog
	engine.Vin = req.Vin
	engine.Date = ToModelDate(req.MsgDate)
	engine.CreatedTime = time.Now()
	err = unmarshalEngine(req.Data, &engine)
	if err != nil {
		return false, err
	}
	engineLog.Engine = engine

	err = repo.EngineUpsert(ctx, engine)
	if err != nil {
		return false, err
	}
	return false, repo.EngineLogCreate(ctx, engineLog)
}

// unmarshalEngine 解析数据
func unmarshalEngine(data []byte, engine *model.Engine) error {
	if len(data) != 5 {
		return fmt.Errorf("发动机数据格式错误:%v", data)
	}
	engine.Status = int32(data[0])
	engine.CrankshaftSpeed = int32(binary.BigEndian.Uint16(data[1:3]))
	engine.FuelConsumption = int32(binary.BigEndian.Uint16(data[3:5]))
	return nil
}
