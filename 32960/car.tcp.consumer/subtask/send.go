package subtask

import (
	"context"
	"fmt"
	"log"

	"car.tcp.consumer/conf"
	"car.tcp.consumer/repo"
	"car.tcp.consumer/tools/httpx"
	"github.com/dapr/go-sdk/service/common"
)

// Command 下行指令
func Command(s common.Service) {
	var CommandSubscription = &common.Subscription{
		PubsubName: conf.Conf.Dapr.PubsubName,
		Topic:      "Command",
		Route:      "/Command",
	}
	if err := s.AddTopicEventHandler(CommandSubscription, commandHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}

// SendMsg .
type SendMsg struct {
	Vin string `json:"vin"`
	Cmd string `json:"cmd"`
}

// Command 指令
func commandHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	var msg SendMsg
	vconn, err := repo.GetTCPAddr(ctx, msg.Vin)
	if err != nil {
		return false, err
	} else if vconn.Host == "" {
		return true, nil
	}

	host := fmt.Sprintf("%s:%s", vconn.Host, vconn.Port)
	var resp struct {
		Code string `json:"code"`
		Data string `json:"data"`
	}
	err = httpx.PostJSON(ctx, httpx.RequestURL(host, "cmdDown"), msg.Cmd, &resp)
	if err != nil || resp.Code != "200" {
		log.Printf("resp:%v, err:%v\n", resp, err)
		return true, fmt.Errorf("resp:%v, err:%v", resp.Data, err.Error())
	}

	return false, nil
}
