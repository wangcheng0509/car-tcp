package base

import (
	"context"
	"fmt"

	"car.tcp.service/conf"
	"car.tcp.service/entity/handleEntity"
	"car.tcp.service/repository"

	"github.com/aceld/zinx/ziface"
	dapr "github.com/dapr/go-sdk/client"
)

var TcpServer ziface.IServer
var DaprClient dapr.Client

func DaprPub(topicName string, publishEventData interface{}) error {
	ctx := context.Background()
	fmt.Println(topicName)
	if err := DaprClient.PublishEvent(ctx, conf.Conf.Dapr.PubsubName, topicName, publishEventData); err != nil {
		return err
	}
	return nil
}

func DoConnectionLost(conn ziface.IConnection) {
	connIdStr := conn.GetConnIdStr()
	repository.OfflineByConnIdStr(conf.Conf.Local.LocalHost, connIdStr)
}

func GetMessage(data []byte) handleEntity.Message {
	msg := handleEntity.GetMessage(data)
	repository.MsgLogAdd(msg)
	return msg
}
