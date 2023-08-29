package subtasktest

import (
	"context"
	"encoding/binary"

	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/subtask"
)

// VlogoutTest ..
func VlogoutTest() {
	data := vlogoutbytes()
	req := msgModel.Message{
		Head:     "##",
		CmdCFlag: 1,
		CmdRsp:   0xFE,
		Vin:      "43302201001018131",
		Encrypt:  0x01,
		Len:      uint16(len(data)),
		Data:     data,
	}

	subtask.VlogoutTest(context.Background(), req)
}

func vlogoutbytes() (data []byte) {
	data = append(data, []byte{23, 8, 14, 14, 44, 30}...)
	data = binary.BigEndian.AppendUint16(data, 12345)
	return
}
