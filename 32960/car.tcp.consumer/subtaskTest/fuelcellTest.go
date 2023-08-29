package subtasktest

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"car.tcp.consumer/entity/msgModel"
	"car.tcp.consumer/subtask"
	"car.tcp.consumer/tools"
)

func FuelcellTest() {
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
	reqParam := fuelcellModel{
		CellCurrent:       100,
		CellVoltage:       220,
		DcStatus:          0x02,
		FuelConsumption:   500,
		H_MaxConc:         500,
		H_ConcSensorCode:  33,
		H_MaxPress:        444,
		H_PressSensorCode: 44,
		H_MaxTemp:         1423,
		H_TempProbeCode:   55,
		ProbeNum:          3,
		ProbeTemps:        "20,30,40",
	}
	reqByte := reqParam.Marshal()
	req.Data = reqByte
	subtask.FuelcellTest(req)
}

type fuelcellModel struct {
	CellCurrent       int       `json:"cellCurrent" gorm:"column:cellCurrent"`             // 燃料电池电流
	CellVoltage       int       `json:"cellVoltage" gorm:"column:cellVoltage"`             // 燃料电池电压
	DcStatus          int       `json:"dcStatus" gorm:"column:dcStatus"`                   // 高压DC/DC状态
	FuelConsumption   int       `json:"fuelConsumption" gorm:"column:fuelConsumption"`     // 燃料消耗率
	H_MaxConc         int       `json:"h_MaxConc" gorm:"column:h_MaxConc"`                 // 氢气最高浓度
	H_ConcSensorCode  int       `json:"h_ConcSensorCode" gorm:"column:h_ConcSensorCode"`   // 氢气最高浓度传感器代号
	H_MaxPress        int       `json:"h_MaxPress" gorm:"column:h_MaxPress"`               // 氢气最高压力
	H_PressSensorCode int       `json:"h_PressSensorCode" gorm:"column:h_PressSensorCode"` // 氢气最高压力传感器代号
	H_MaxTemp         int       `json:"h_MaxTemp" gorm:"column:h_MaxTemp"`                 // 氢系统中最高温度
	H_TempProbeCode   int       `json:"h_TempProbeCode" gorm:"column:h_TempProbeCode"`     // 氢系统中最高温度探针代号
	ProbeNum          int       `json:"probeNum" gorm:"column:probeNum"`                   // 燃料电池温度探针总数
	ProbeTemps        string    `json:"probeTemps" gorm:"column:probeTemps"`               // 探针温度值
	CreatedTime       time.Time `json:"created_time" gorm:"column:created_time"`           // 上报时间
}

func (v fuelcellModel) Marshal() []byte {
	var bytes []byte
	CellVoltagebytes, err := tools.GetByteFromUint16(int(v.CellVoltage))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, CellVoltagebytes...)
	}
	CellCurrentbytes, err := tools.GetByteFromUint16(int(v.CellCurrent))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, CellCurrentbytes...)
	}
	FuelConsumptionbytes, err := tools.GetByteFromUint16(int(v.FuelConsumption))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, FuelConsumptionbytes...)
	}
	ProbeNumbytes, err := tools.GetByteFromUint16(int(v.ProbeNum))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, ProbeNumbytes...)
	}
	probeTempsList := strings.Split(v.ProbeTemps, ",")
	for _, v := range probeTempsList {
		vi, _ := strconv.Atoi(v)
		bytes = append(bytes, byte(vi))
	}
	HMaxTempbytes, err := tools.GetByteFromUint16(int(v.H_MaxTemp))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, HMaxTempbytes...)
	}
	bytes = append(bytes, byte(v.H_TempProbeCode))
	HMaxConcbytes, err := tools.GetByteFromUint16(int(v.H_MaxConc))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, HMaxConcbytes...)
	}
	bytes = append(bytes, byte(v.H_ConcSensorCode))
	HMaxPressbytes, err := tools.GetByteFromUint16(int(v.H_MaxPress))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, HMaxPressbytes...)
	}
	bytes = append(bytes, byte(v.H_PressSensorCode))
	bytes = append(bytes, byte(v.DcStatus))

	fmt.Println("燃料电池数据", len(bytes))
	fmt.Println("燃料电池数据", bytes)
	return bytes

}
