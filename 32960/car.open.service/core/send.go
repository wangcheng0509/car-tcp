package core

import (
	"context"
	"fmt"

	"car.open.service/conf"
	dapr "github.com/dapr/go-sdk/client"
)

// DaprClient ..
var DaprClient dapr.Client

// SendMsg .
type SendMsg struct {
	Vin string `json:"vin"`
	Cmd string `json:"cmd"`
}

// Command 指令
func Command(ctx context.Context, vin string, cmd string) error {
	if err := DaprClient.PublishEvent(ctx,
		conf.Conf.Dapr.PubsubName,
		"Commmand",
		SendMsg{
			Vin: vin,
			Cmd: cmd,
		},
	); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
