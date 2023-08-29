package subtasktest

import (
	"fmt"

	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/subtask"
	"car.tcp.consumer/tools"
)

func VehicleTest() {
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
	reqParam := vehicleModel{
		AcceleratorPedal: 1,
		BrakePedal:       1,
		Charging:         1,
		Current:          11000,
		Dc:               2,
		Gear:             62,
		Mileage:          122222,
		Mode:             1,
		Resistance:       51236,
		Soc:              50,
		Speed:            0,
		Status:           2,
		Voltage:          400,
	}
	reqByte := reqParam.Marshal()
	req.Data = reqByte
	subtask.VehicleTest(req)
}

type vehicleModel struct {
	AcceleratorPedal int
	BrakePedal       int
	Charging         int
	Current          int
	Dc               int
	Gear             int
	Mileage          int
	Mode             int
	Resistance       int
	Soc              int
	Speed            int
	Status           int
	Voltage          int
}

func (v vehicleModel) Marshal() []byte {
	var bytes []byte
	bytes = append(bytes, byte(v.Status))
	bytes = append(bytes, byte(v.Charging))
	bytes = append(bytes, byte(v.Mode))
	speedbytes, err := tools.GetByteFromUint16(int(v.Speed))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, speedbytes...)
	}
	mileagebytes, err := tools.GetByteFromUint32(int(v.Mileage))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, mileagebytes...)
	}
	voltagebytes, err := tools.GetByteFromUint16(int(v.Voltage))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, voltagebytes...)
	}
	currentbytes, err := tools.GetByteFromUint16(int(v.Current))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, currentbytes...)
	}
	bytes = append(bytes, byte(v.Soc))
	bytes = append(bytes, byte(v.Dc))
	bytes = append(bytes, byte(v.Gear))
	resistancebytes, err := tools.GetByteFromUint16(int(v.Resistance))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, resistancebytes...)
	}
	bytes = append(bytes, byte(v.BrakePedal))
	bytes = append(bytes, byte(v.AcceleratorPedal))
	fmt.Println("整车信息", len(bytes))
	fmt.Println("整车信息", bytes)
	return bytes

}
