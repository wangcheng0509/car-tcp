package subtasktest

import (
	"context"
	"encoding/binary"

	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/subtask"
)

// EngineTest ..
func EngineTest() {
	req := msgModel.InfoCateModel{
		Vin:     "43302201001018131",
		CmdFlag: 2,
		MsgDate: msgModel.MsgDate{
			Year:    23,
			Month:   8,
			Day:     14,
			Hour:    14,
			Minutes: 44,
			Seconds: 30,
		},
	}

	req.Data = enginebytes()
	subtask.EngineTest(context.Background(), req)
}

func enginebytes() (data []byte) {
	data = append(data, 0x01)
	data = binary.BigEndian.AppendUint16(data, 2000)
	data = binary.BigEndian.AppendUint16(data, 500)
	return
}
