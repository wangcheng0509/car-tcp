package subtask

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"car.tcp.consumer/app"
	"car.tcp.consumer/conf"
	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/tools"
	"github.com/dapr/go-sdk/service/common"
	"github.com/wangcheng0509/gpkg/try"
)

func RealtimeInfo(s common.Service) {
	var realtimeInfoSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "RealtimeInfo",
		Route:      "/RealtimeInfo",
	}
	if err := s.AddTopicEventHandler(realtimeInfoSubscription, realtimeInfoHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

func RealtimeInfoTest(req msgModel.Message) {
	// 32960协议解析
	infoUnmarshal(req)
}

func realtimeInfoHandler(ctx context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	// 解析请求msg
	var req msgModel.Message
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		fmt.Println("解析请求msg Error:", err.Error())
		return false, err
	}
	fmt.Println("msg解析:", req)
	// 32960协议解析
	err := infoUnmarshal(req)
	if err != nil {
		return true, err
	}
	return false, nil
}

func infoUnmarshal(req msgModel.Message) error {
	var err error
	var msgDate msgModel.MsgDate
	if len(req.Data) < 7 {
		return errors.New("param is error")
	}
	try.Try(func() {
		defer func() {
			if errInner := recover(); errInner != nil {
				try.Throw(1, errInner.(string))
			}
		}()
		// 年月日时分秒
		if msgDate, err = tools.GetMsgDate(req.Data); err != nil {
			try.Throw(1, "param is error")
		}
		index := 6
		for {
			if index >= len(req.Data)-1 {
				break
			}
			cmdFlg := req.Data[index]
			index = index + 1
			switch {
			case cmdFlg == 0x01:
				// 整车数据
				app.DaprPub("Vehicle", msgModel.InfoCateModel{
					Vin:     req.Vin,
					CmdFlag: cmdFlg,
					MsgDate: msgDate,
					Data:    req.Data[index : index+20],
				})
				index = index + 20
			case cmdFlg == 0x02:
				// 驱动电机数据
				num := int(req.Data[index])
				app.DaprPub("Drivemotor", msgModel.InfoCateModel{
					Vin:     req.Vin,
					CmdFlag: cmdFlg,
					MsgDate: msgDate,
					Data:    req.Data[index : index+1+(num*12)],
				})
				index = index + 1 + (num * 12)
			case cmdFlg == 0x03:
				// 燃料电池数据
				num := tools.GetUInt16((req.Data[index+6 : index+6+2]))
				app.DaprPub("Fuelcell", msgModel.InfoCateModel{
					Vin:     req.Vin,
					CmdFlag: cmdFlg,
					MsgDate: msgDate,
					Data:    req.Data[index : index+18+num],
				})
				index = index + 18 + num
			case cmdFlg == 0x04:
				// 发动机数据
				app.DaprPub("Engine", msgModel.InfoCateModel{
					Vin:     req.Vin,
					CmdFlag: cmdFlg,
					MsgDate: msgDate,
					Data:    req.Data[index : index+5],
				})
				index = index + 5
			case cmdFlg == 0x05:
				// 车辆位置数据
				app.DaprPub("Location", msgModel.InfoCateModel{
					Vin:     req.Vin,
					CmdFlag: cmdFlg,
					MsgDate: msgDate,
					Data:    req.Data[index : index+9],
				})
				index = index + 9
			case cmdFlg == 0x06:
				// 极值数据
				app.DaprPub("Extreme", msgModel.InfoCateModel{
					Vin:     req.Vin,
					CmdFlag: cmdFlg,
					MsgDate: msgDate,
					Data:    req.Data[index : index+14],
				})
				index = index + 14
			case cmdFlg == 0x07:
				// 报警数据
				// // 可充电储能装置故障总数N1
				// num1 := tools.GetUInt16((req.Data[index+5 : index+5+1]))
				// // 驱动电机故障总数N2
				// num2 := tools.GetUInt16((req.Data[index+6 : index+6+1]))
				// // 发动机故障总数N3
				// num3 := tools.GetUInt16((req.Data[index+6 : index+6+1]))
				// // 其他故障总数N4
				// num4 := tools.GetUInt16((req.Data[index+6 : index+6+1]))

			case cmdFlg == 0x08:
				fmt.Println(req.Data, "index:", index)
				num := int(req.Data[index])
				m := int(req.Data[index+10])
				app.DaprPub("ChargeableVoltage", msgModel.InfoCateModel{
					Vin:     req.Vin,
					CmdFlag: cmdFlg,
					MsgDate: msgDate,
					Data:    req.Data[index : index+num*(10+2*m)+1],
				})
				index = index + num*(10+2*m)

			case cmdFlg == 0x09:
				fmt.Println(req.Data, "index:", index)
				num := int(req.Data[index])
				m := int(tools.GetUInt16(req.Data[index+2 : index+4]))
				app.DaprPub("ChargeableTemp", msgModel.InfoCateModel{
					Vin:     req.Vin,
					CmdFlag: cmdFlg,
					MsgDate: msgDate,
					Data:    req.Data[index : index+num*(3+m)+1],
				})
				index = index + num*(3+m)
			default:
				break
			}

		}
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return err
}
