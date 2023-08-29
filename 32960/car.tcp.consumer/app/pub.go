package app

import (
	"context"
	"fmt"

	"car.tcp.consumer/conf"
	dapr "github.com/dapr/go-sdk/client"
)

var DaprClient dapr.Client

func DaprPub(topicName string, publishEventData interface{}) error {
	ctx := context.Background()
	if err := DaprClient.PublishEvent(ctx, conf.Conf.Dapr.PubsubName, topicName, publishEventData); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
