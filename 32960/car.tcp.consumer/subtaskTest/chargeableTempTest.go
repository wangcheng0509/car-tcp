package subtasktest

import (
	"context"
	"encoding/binary"

	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/subtask"
)

// ChargeableTempTest ..
func ChargeableTempTest() {
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

	req.Data = chargeableTempbytes()
	subtask.ChargeableTempTest(context.Background(), req)
}

func chargeableTempbytes() (data []byte) {
	data = append(data, 0x02)

	data = append(data, 1)
	data = binary.BigEndian.AppendUint16(data, 2)
	data = append(data, 50)
	data = append(data, 60)

	data = append(data, 2)
	data = binary.BigEndian.AppendUint16(data, 2)
	data = append(data, 60)
	data = append(data, 65)

	return
}
