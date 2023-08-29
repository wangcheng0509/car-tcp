package subtasktest

import (
	"context"
	"encoding/binary"

	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/subtask"
)

// DrivemotorTest ..
func DrivemotorTest() {
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

	req.Data = drivemotorbytes()
	subtask.DrivemotorTest(context.Background(), req)
}

func drivemotorbytes() (data []byte) {
	data = append(data, 2)
	data = append(data, 1)
	data = append(data, 0x01)
	data = append(data, 80)
	data = binary.BigEndian.AppendUint16(data, 40000)
	data = binary.BigEndian.AppendUint16(data, 50000)
	data = append(data, 60)
	data = binary.BigEndian.AppendUint16(data, 2000)
	data = binary.BigEndian.AppendUint16(data, 1020)
	data = append(data, 2)
	data = append(data, 0x02)
	data = append(data, 180)
	data = binary.BigEndian.AppendUint16(data, 20000)
	data = binary.BigEndian.AppendUint16(data, 30000)
	data = append(data, 70)
	data = binary.BigEndian.AppendUint16(data, 2100)
	data = binary.BigEndian.AppendUint16(data, 1080)

	return
}
