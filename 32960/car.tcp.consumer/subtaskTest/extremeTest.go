package subtasktest

import (
	"fmt"

	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/subtask"
	"car.tcp.consumer/tools"
)

func ExtremeTest() {
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
	reqParam := extreme{
		MaxBatteryVoltage:         300,
		MaxTemp:                   150,
		MaxTempProbeNo:            11,
		MaxTempSubsysNo:           111,
		MaxVoltageBatteryCode:     22,
		MaxVoltageBatterySubsysNo: 222,
		MinBatteryVoltage:         220,
		MinTemp:                   100,
		MinTempProbeNo:            33,
		MinTempSubsysNo:           333,
		MinVoltageBatteryCode:     44,
		MinVoltageBatterySubsysNo: 444,
	}
	reqByte := reqParam.Marshal()
	req.Data = reqByte
	subtask.ExtremeTest(req)
}

type extremeModel struct {
	MaxBatteryVoltage         float32 `json:"maxBatteryVoltage" gorm:"column:maxBatteryVoltage"`                 // 电池单体电压最高值
	MaxTemp                   int     `json:"maxTemp" gorm:"column:maxTemp"`                                     // 最高温度值
	MaxTempProbeNo            int     `json:"maxTempProbeNo" gorm:"column:maxTempProbeNo"`                       // 最高温度探针序号
	MaxTempSubsysNo           int     `json:"maxTempSubsysNo" gorm:"column:maxTempSubsysNo"`                     // 最高温度子系统号
	MaxVoltageBatteryCode     int     `json:"maxVoltageBatteryCode" gorm:"column:maxVoltageBatteryCode"`         // 最高电压电池单体代号
	MaxVoltageBatterySubsysNo int     `json:"maxVoltageBatterySubsysNo" gorm:"column:maxVoltageBatterySubsysNo"` // 最高电压电池子系统号
	MinBatteryVoltage         float32 `json:"minBatteryVoltage" gorm:"column:minBatteryVoltage"`                 // 电池单体电压最低值
	MinTemp                   int     `json:"minTemp" gorm:"column:minTemp"`                                     // 最低温度值
	MinTempProbeNo            int     `json:"minTempProbeNo" gorm:"column:minTempProbeNo"`                       // 最低温度探针序号
	MinTempSubsysNo           int     `json:"minTempSubsysNo" gorm:"column:minTempSubsysNo"`                     // 最低温度子系统号
	MinVoltageBatteryCode     int     `json:"minVoltageBatteryCode" gorm:"column:minVoltageBatteryCode"`         // 最低电压电池单体代号
	MinVoltageBatterySubsysNo int     `json:"minVoltageBatterySubsysNo" gorm:"column:minVoltageBatterySubsysNo"` // 最低电压电池子系统号
}

func (v extremeModel) Marshal() []byte {
	var bytes []byte
	bytes = append(bytes, byte(v.MaxVoltageBatterySubsysNo))
	bytes = append(bytes, byte(v.MaxVoltageBatteryCode))
	MaxBatteryVoltagebytes, err := tools.GetByteFromUint16(int(v.MaxBatteryVoltage))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, MaxBatteryVoltagebytes...)
	}
	bytes = append(bytes, byte(v.MinVoltageBatterySubsysNo))
	bytes = append(bytes, byte(v.MinVoltageBatteryCode))
	MinBatteryVoltagebytes, err := tools.GetByteFromUint16(int(v.MinBatteryVoltage))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, MinBatteryVoltagebytes...)
	}

	bytes = append(bytes, byte(v.MaxTempSubsysNo))
	bytes = append(bytes, byte(v.MaxTempProbeNo))
	bytes = append(bytes, byte(v.MaxTemp))
	bytes = append(bytes, byte(v.MinTempSubsysNo))
	bytes = append(bytes, byte(v.MinTempProbeNo))
	bytes = append(bytes, byte(v.MinTemp))
	fmt.Println("极值数据", len(bytes))
	fmt.Println("极值数据", bytes)
	return bytes
}
