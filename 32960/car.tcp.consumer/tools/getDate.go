package tools

import (
	"errors"

	"car.tcp.consumer/entity/msgModel"
	"github.com/wangcheng0509/gpkg/try"
)

func GetMsgDate(data []byte) (msgModel.MsgDate, error) {
	var err error
	var msgDate msgModel.MsgDate
	if len(data) < 6 {
		return msgDate, errors.New("param is error")
	}
	try.Try(func() {
		defer func() {
			if errInner := recover(); errInner != nil {
				try.Throw(1, errInner.(string))
			}
		}()
		msgDate.Year = int(data[0])
		msgDate.Month = int(data[1])
		msgDate.Day = int(data[2])
		msgDate.Hour = int(data[3])
		msgDate.Minutes = int(data[4])
		msgDate.Seconds = int(data[5])
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})
	return msgDate, err
}
