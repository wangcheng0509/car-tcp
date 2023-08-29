package tcpclient

import (
	"client/tools"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

func getinfoByte() []byte {
	reqParam := dateModel{
		Year:    23,
		Month:   8,
		Day:     16,
		Hour:    17,
		Minutes: 17,
		Seconds: 30,
	}
	reqByte := reqParam.Marshal()
	// 整车信息
	reqByte = append(reqByte, byte(0x01))
	vehicleParam := vehicle{
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
		Voltage:          220,
	}
	reqByte = append(reqByte, vehicleParam.Marshal()...)

	// 驱动电机数据
	reqByte = append(reqByte, byte(0x02))
	reqByte = append(reqByte, byte(2))
	drivemotorParam1 := drivemotor{
		No:           1,
		CtrlTemp:     3,
		DcBusCurrent: 11000,
		InputVoltage: 220,
		MotorTemp:    150,
		Rotating:     50000,
		Status:       3,
		Torque:       50000,
	}
	reqByte = append(reqByte, drivemotorParam1.Marshal()...)
	drivemotorParam2 := drivemotor{
		No:           1,
		CtrlTemp:     3,
		DcBusCurrent: 11000,
		InputVoltage: 220,
		MotorTemp:    150,
		Rotating:     50000,
		Status:       3,
		Torque:       50000,
	}
	reqByte = append(reqByte, drivemotorParam2.Marshal()...)
	// // 燃料电池数据
	reqByte = append(reqByte, byte(0x03))
	fuelcellParam := fuelcell{
		CellCurrent:       10000,
		CellVoltage:       220,
		DcStatus:          0x02,
		FuelConsumption:   50000,
		H_MaxConc:         50000,
		H_ConcSensorCode:  33,
		H_MaxPress:        444,
		H_PressSensorCode: 44,
		H_MaxTemp:         1423,
		H_TempProbeCode:   55,
		ProbeNum:          3,
		ProbeTemps:        "20,30,40",
	}
	reqByte = append(reqByte, fuelcellParam.Marshal()...)

	// 发动机数据
	reqByte = append(reqByte, byte(0x04))
	engineParam := engine{
		CrankshaftSpeed: 50000,
		FuelConsumption: 20000,
		Status:          0x02,
	}
	reqByte = append(reqByte, engineParam.Marshal()...)

	// 车辆位置数据
	reqByte = append(reqByte, byte(0x05))
	locationParam := location{
		Latitude:  3664729,
		Longitude: 11708357,
		Status:    32,
	}
	reqByte = append(reqByte, locationParam.Marshal()...)
	// 极值数据
	reqByte = append(reqByte, byte(0x06))
	extremeParam := extreme{
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
	reqByte = append(reqByte, extremeParam.Marshal()...)

	// 可充电储能装置电压数据
	reqByte = append(reqByte, byte(0x08))
	reqByte = append(reqByte, chargeableVoltagebytes()...)
	// // 可充电储能装置温度数据
	reqByte = append(reqByte, byte(0x09))
	reqByte = append(reqByte, chargeableTempbytes()...)

	return reqByte
}

// 日期
type dateModel struct {
	Year    int `json:"year" gorm:"column:year"`
	Month   int `json:"month" gorm:"column:month"`
	Day     int `json:"day" gorm:"column:day"`
	Hour    int `json:"hour" gorm:"column:hour"`
	Minutes int `json:"minutes" gorm:"column:minutes"`
	Seconds int `json:"seconds" gorm:"column:seconds"`
}

func (v dateModel) Marshal() []byte {
	var bytes []byte
	bytes = append(bytes, byte(v.Year))
	bytes = append(bytes, byte(v.Month))
	bytes = append(bytes, byte(v.Day))
	bytes = append(bytes, byte(v.Hour))
	bytes = append(bytes, byte(v.Minutes))
	bytes = append(bytes, byte(v.Seconds))
	return bytes
}

// 整车信息
type vehicle struct {
	AcceleratorPedal int `json:"acceleratorPedal" gorm:"column:acceleratorPedal"`
	BrakePedal       int `json:"brakePedal" gorm:"column:brakePedal"`
	Charging         int `json:"charging" gorm:"column:charging"`     // 充电状态；0x01;停车充电;0x02;行驶充电;0x03;未充电状态;x04:充电完成;“0xFE”表示异常，“0xFE”表示无效
	Current          int `json:"current" gorm:"column:current"`       // 总电流；有效值范围:0~20 000(偏移量1 000 A，表示一1000 A~+1 000 A)最小计量单元01A，“0xFFOXFE”表示异常“0xFF0xFF”表示无效
	Dc               int `json:"dc" gorm:"column:dc"`                 // 0x01：工作0x02：断开，"OxFE”表示异常，"OxFF”表示无效
	Gear             int `json:"gear" gorm:"column:gear"`             // 挡位；
	Mileage          int `json:"mileage" gorm:"column:mileage"`       // 累计里程；有效值范围:0~9 999 999(表示0 km~999 999.9 km)最小计量单元;0.1 km“0xFF，0xFF，0xFF,OxFE”表示异常，“0xFF,OxFF,OxFF0xFF”示无效
	Mode             int `json:"mode" gorm:"column:mode"`             // 运行模式；0x01:纯电;0x02;混动;0x03:燃油;0xFE表示异常;0xFF 表示无效
	Resistance       int `json:"resistance" gorm:"column:resistance"` // 绝缘电阻；有效范围0~60 000(表示 0 Q~60 000 k2)最小计量单元;1 k2
	Soc              int `json:"soc" gorm:"column:soc"`               // 有效值范围;0~100(表示 0%~100%),最小计量单元:1%“0xFE”表示异常“0xFF”表示无效
	Speed            int `json:"speed" gorm:"column:speed"`           // 车速；有效值范围;0~100(表示 0%~100%),最小计量单元:1%，“0xFE”表示异常“0xFF”表示无效
	Status           int `json:"status" gorm:"column:status"`         // 车辆状态；0x01:启动，0x02:熄火，0x03:其他，0xFE表示异常，0xFF表示无效
	Voltage          int `json:"voltage" gorm:"column:voltage"`       // 总电压；有效值范围:0~10 000(表示 0 V~1000 V)，小计量单元:0.1V，“0xFF,xFE”示异常，“0xFFxFF”表示无效
}

func (v vehicle) Marshal() []byte {
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

// 驱动电机数据
type drivemotor struct {
	No           int `json:"no" gorm:"column:no"`                     // 驱动电机序号
	CtrlTemp     int `json:"ctrlTemp" gorm:"column:ctrlTemp"`         // 驱动电机控制器温度
	DcBusCurrent int `json:"dcBusCurrent" gorm:"column:dcBusCurrent"` // 电机控制器直流母线电流
	InputVoltage int `json:"inputVoltage" gorm:"column:inputVoltage"` // 电机控制器输人电压
	MotorTemp    int `json:"motorTemp" gorm:"column:motorTemp"`       // 驱动电机温度
	Rotating     int `json:"rotating" gorm:"column:rotating"`         // 驱动电机转速
	Status       int `json:"status" gorm:"column:status"`             // 驱动电机状态
	Torque       int `json:"torque" gorm:"column:torque"`             // 驱动电机转矩
}

func (v drivemotor) Marshal() []byte {
	var bytes []byte
	bytes = append(bytes, byte(v.No))
	bytes = append(bytes, byte(v.Status))
	bytes = append(bytes, byte(v.CtrlTemp))
	Rotatingbytes, err := tools.GetByteFromUint16(int(v.Rotating))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, Rotatingbytes...)
	}
	Torquebytes, err := tools.GetByteFromUint16(int(v.Torque))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, Torquebytes...)
	}
	bytes = append(bytes, byte(v.MotorTemp))
	InputVoltagebytes, err := tools.GetByteFromUint16(int(v.InputVoltage))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, InputVoltagebytes...)
	}
	DcBusCurrentbytes, err := tools.GetByteFromUint16(int(v.DcBusCurrent))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, DcBusCurrentbytes...)
	}
	fmt.Println("驱动电机数据", len(bytes))
	fmt.Println("驱动电机数据", bytes)
	return bytes
}

// 燃料电池数据
type fuelcell struct {
	CellCurrent       int    `json:"cellCurrent" gorm:"column:cellCurrent"`             // 燃料电池电流
	CellVoltage       int    `json:"cellVoltage" gorm:"column:cellVoltage"`             // 燃料电池电压
	DcStatus          int    `json:"dcStatus" gorm:"column:dcStatus"`                   // 高压DC/DC状态
	FuelConsumption   int    `json:"fuelConsumption" gorm:"column:fuelConsumption"`     // 燃料消耗率
	H_MaxConc         int    `json:"h_MaxConc" gorm:"column:h_MaxConc"`                 // 氢气最高浓度
	H_ConcSensorCode  int    `json:"h_ConcSensorCode" gorm:"column:h_ConcSensorCode"`   // 氢气最高浓度传感器代号
	H_MaxPress        int    `json:"h_MaxPress" gorm:"column:h_MaxPress"`               // 氢气最高压力
	H_PressSensorCode int    `json:"h_PressSensorCode" gorm:"column:h_PressSensorCode"` // 氢气最高压力传感器代号
	H_MaxTemp         int    `json:"h_MaxTemp" gorm:"column:h_MaxTemp"`                 // 氢系统中最高温度
	H_TempProbeCode   int    `json:"h_TempProbeCode" gorm:"column:h_TempProbeCode"`     // 氢系统中最高温度探针代号
	ProbeNum          int    `json:"probeNum" gorm:"column:probeNum"`                   // 燃料电池温度探针总数
	ProbeTemps        string `json:"probeTemps" gorm:"column:probeTemps"`               // 探针温度值
}

func (v fuelcell) Marshal() []byte {
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

// 发动机数据
type engine struct {
	CrankshaftSpeed int     `json:"crankshaftSpeed" gorm:"column:crankshaftSpeed"` // 曲轴转速；有效范围：0-60 000（表示0 r/min-60 000 r/min)，最小计量单元：1 r/min,0xFF/0xFE表示异常，0xFF,0xFF表示无效
	FuelConsumption float32 `json:"fuelConsumption" gorm:"column:fuelConsumption"` // 燃料消耗率；有效范围:0~60 000(表示0r/min一60 000 r/min)，最小计量单元:1r/min，0xFF/OxFE表示异常，0xFF表示无效
	Status          int     `json:"status" gorm:"column:status"`                   // 发动机状态；0x01启动状态；0x02关闭状态，0xFE表示异常，0xFF表示无效
}

func (v engine) Marshal() []byte {
	var bytes []byte
	bytes = append(bytes, byte(v.Status))
	CrankshaftSpeedbytes, err := tools.GetByteFromUint16(int(v.CrankshaftSpeed))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, CrankshaftSpeedbytes...)
	}
	FuelConsumptionbytes, err := tools.GetByteFromUint16(int(v.FuelConsumption))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, FuelConsumptionbytes...)
	}
	fmt.Println("发动机数据", len(bytes))
	fmt.Println("发动机数据", bytes)
	return bytes
}

// 车辆位置数据
type location struct {
	Latitude  int64 `json:"latitude" gorm:"column:latitude"`   // 纬度；以度为单位的纬度值乘以10^6，精确到百万分之一度
	Longitude int64 `json:"longitude" gorm:"column:longitude"` // 经度；以度为单位的纬度值乘以10^6，精确到百万分之一度
	Status    int   `json:"status" gorm:"column:status"`
}

func (v location) Marshal() []byte {
	var bytes []byte
	bytes = append(bytes, byte(v.Status))
	Longitudebytes, err := tools.GetByteFromUint32(int(v.Longitude))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, Longitudebytes...)
	}
	Latitudebytes, err := tools.GetByteFromUint32(int(v.Latitude))
	if err != nil {
		panic(err)
	} else {
		bytes = append(bytes, Latitudebytes...)
	}
	fmt.Println("车辆位置数据", len(bytes))
	fmt.Println("车辆位置数据", bytes)
	return bytes
}

// 极值数据
type extreme struct {
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

func (v extreme) Marshal() []byte {
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

func chargeableVoltagebytes() (data []byte) {
	// 可充电储能子系统个数
	data = append(data, 2)

	// 可充电储能子系统电压信息
	data = append(data, 1)
	data = binary.BigEndian.AppendUint16(data, 200)
	data = binary.BigEndian.AppendUint16(data, 1020)
	data = binary.BigEndian.AppendUint16(data, 2)
	data = binary.BigEndian.AppendUint16(data, 1)
	data = append(data, 2)
	data = binary.BigEndian.AppendUint16(data, 20000)
	data = binary.BigEndian.AppendUint16(data, 20000)

	// 可充电储能子系统电压信息
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

func chargeableTempbytes() (data []byte) {
	// 可充电储能子系统个数
	data = append(data, 2)

	// 可充电储能子系统温度信息
	data = append(data, 1)
	data = binary.BigEndian.AppendUint16(data, 2)
	data = append(data, 50)
	data = append(data, 60)

	// 可充电储能子系统温度信息
	data = append(data, 2)
	data = binary.BigEndian.AppendUint16(data, 2)
	data = append(data, 60)
	data = append(data, 65)

	return
}
