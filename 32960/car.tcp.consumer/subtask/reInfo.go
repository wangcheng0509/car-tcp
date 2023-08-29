package subtask

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"car.tcp.consumer/conf"
	"car.tcp.consumer/entity/msgModel"
	"github.com/dapr/go-sdk/service/common"
)

func ResendInfo(s common.Service) {
	var ResendInfoSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "ResendInfo",
		Route:      "/ResendInfo",
	}
	if err := s.AddTopicEventHandler(ResendInfoSubscription, resendInfoHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

func resendInfoHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	var req msgModel.Message
	if err := json.Unmarshal(e.RawData, &req); err != nil {
		return false, err
	}
	fmt.Println(req)
	return false, nil
}
