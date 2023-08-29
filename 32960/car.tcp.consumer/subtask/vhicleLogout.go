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
	"car.tcp.consumer/tools"
	"github.com/dapr/go-sdk/service/common"
)

func VehicleLogout(s common.Service) {
	var VehicleLogoutSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "VehicleLogout",
		Route:      "/VehicleLogout",
	}
	if err := s.AddTopicEventHandler(VehicleLogoutSubscription, vehicleLogoutHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

// VlogoutTest ..
func VlogoutTest(ctx context.Context, req msgModel.Message) (retry bool, err error) {
	var vlogout model.Vlogout
	var vlogoutLog model.VlogoutLog
	err = unmarshalvehicleLogout(req.Data, &vlogout)
	if err != nil {
		return false, err
	}
	vlogout.Vin = req.Vin
	vlogout.CreatedTime = time.Now()
	vlogoutLog.Vlogout = vlogout

	err = repo.NewTransRepo().Exec(ctx, func(c context.Context) error {
		err := repo.VlogoutUpsert(c, vlogout)
		if err != nil {
			return err
		}
		return repo.VlogoutLogCreate(c, vlogoutLog)
	})
	if err != nil {
		return false, err
	}

	return false, nil
}

func vehicleLogoutHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	var req msgModel.Message
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		return false, err
	}
	var vlogout model.Vlogout
	var vlogoutLog model.VlogoutLog
	err = unmarshalvehicleLogout(req.Data, &vlogout)
	if err != nil {
		return false, err
	}
	vlogout.Vin = req.Vin
	vlogout.CreatedTime = time.Now()
	vlogoutLog.Vlogout = vlogout

	err = repo.NewTransRepo().Exec(ctx, func(c context.Context) error {
		err := repo.VlogoutUpsert(c, vlogout)
		if err != nil {
			return err
		}
		return repo.VlogoutLogCreate(c, vlogoutLog)
	})
	if err != nil {
		return false, err
	}

	return false, nil
}

func unmarshalvehicleLogout(data []byte, vlogout *model.Vlogout) error {
	if len(data) != 8 {
		return fmt.Errorf("车辆登出数据格式错误:%v", data)
	}
	date, err := tools.GetMsgDate(data)
	if err != nil {
		return err
	}
	vlogout.Date = ToModelDate(date)
	vlogout.Seq = int(binary.BigEndian.Uint16(data[6:8]))

	return nil
}
