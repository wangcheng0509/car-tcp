package subtasktest

import (
	"context"
	"encoding/binary"

	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/subtask"
)

// ChargeableVoltageTest ..
func ChargeableVoltageTest() {
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

	req.Data = chargeableVoltagebytes()
	subtask.ChargeableVoltageTest(context.Background(), req)
}

func chargeableVoltagebytes() (data []byte) {
	data = append(data, 2)

	data = append(data, 1)
	data = binary.BigEndian.AppendUint16(data, 200)
	data = binary.BigEndian.AppendUint16(data, 1020)
	data = binary.BigEndian.AppendUint16(data, 2)
	data = binary.BigEndian.AppendUint16(data, 1)
	data = append(data, 2)
	data = binary.BigEndian.AppendUint16(data, 20000)
	data = binary.BigEndian.AppendUint16(data, 20000)

	data = append(data, 2)
	data = binary.BigEndian.AppendUint16(data, 200)
	data = binary.BigEndian.AppendUint16(data, 1020)
	data = binary.BigEndian.AppendUint16(data, 2)
	data = binary.BigEndian.AppendUint16(data, 1)
	data = append(data, 2)
	data = binary.BigEndian.AppendUint16(data, 30000)
	data = binary.BigEndian.AppendUint16(data, 30000)

	return
}
