package subtask

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"car.tcp.consumer/conf"
	"car.tcp.consumer/entity/msgModel"

	"car.tcp.consumer/tools"
	"github.com/dapr/go-sdk/service/common"
	"github.com/wangcheng0509/gpkg/try"
)

// Topic注册
func Heartbeat(s common.Service) {
	var heartbeatSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "Heartbeat",
		Route:      "/Heartbeat",
	}

	if err := s.AddTopicEventHandler(heartbeatSubscription, heartbeatHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

func HeartbeatTest(req msgModel.Message) (retry bool, err error) {
	// 32960协议解析
	v, err := heartbeatUnmarshal(req)
	if err != nil {
		return false, err
	}
	log.Println("Heartbeat:", v)
	// TODO 获取车辆未发送指令，下发指令
	return false, nil
}

// task
func heartbeatHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s", e.PubsubName, e.Topic, e.ID)
	// 解析请求msg
	var req msgModel.Message
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		fmt.Println("解析请求msg Error:", err.Error())
		return false, err
	}
	fmt.Println(req)
	// 32960协议解析
	v, err := heartbeatUnmarshal(req)
	if err != nil {
		fmt.Println("32960协议解析 Error:", err.Error())
		return false, err
	}
	log.Println("Heartbeat:", v)
	// TODO 获取车辆未发送指令，下发指令

	return false, nil
}

func heartbeatUnmarshal(req msgModel.Message) (msgModel.InfoCateModel, error) {
	var msg msgModel.InfoCateModel
	var err error
	if len(req.Data) < 30 {
		return msg, errors.New("param is error")
	}
	try.Try(func() {
		defer func() {
			if errInner := recover(); errInner != nil {
				try.Throw(1, errInner.(string))
			}
		}()
		msg.Vin = req.Vin
		// 年月日时分秒
		if msgDate, err := tools.GetMsgDate(req.Data); err == nil {
			msg.MsgDate.Year = msgDate.Year
			msg.MsgDate.Month = msgDate.Month
			msg.MsgDate.Day = msgDate.Day
			msg.MsgDate.Hour = msgDate.Hour
			msg.MsgDate.Minutes = msgDate.Minutes
			msg.MsgDate.Seconds = msgDate.Seconds
		} else {
			try.Throw(1, "param is error")
		}
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return msg, err
}
