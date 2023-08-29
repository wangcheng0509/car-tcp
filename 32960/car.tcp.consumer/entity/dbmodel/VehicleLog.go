package dbmodel

import "time"

// 整车信息Log
type VehicleLog struct {
	Vin              string    `json:"vin" gorm:"column:vin"`
	Year             int32     `json:"year" gorm:"column:year"`
	Month            int32     `json:"month" gorm:"column:month"`
	Day              int32     `json:"day" gorm:"column:day"`
	Hour             int32     `json:"hour" gorm:"column:hour"`
	Minutes          int32     `json:"minutes" gorm:"column:minutes"`
	Seconds          int32     `json:"seconds" gorm:"column:seconds"`
	AcceleratorPedal int32     `json:"acceleratorPedal" gorm:"column:acceleratorPedal"`
	BrakePedal       int32     `json:"brakePedal" gorm:"column:brakePedal"`
	Charging         int32     `json:"charging" gorm:"column:charging"`         // 充电；0x01;停车充电;0x02;行驶充电;0x03;未充电状态;x04:充电完成;“0xFE”表示异常，“0xFE”表示无效
	Current          int32     `json:"current" gorm:"column:current"`           // 总电流；有效值范围:0~20 000(偏移量1 000 A，表示一1000 A~+1 000 A)最小计量单元01A，“0xFFOXFE”表示异常“0xFF0xFF”表示无效
	Dc               int32     `json:"dc" gorm:"column:dc"`                     // 0x01：工作0x02：断开，"OxFE”表示异常，"OxFF”表示无效
	Gear             int32     `json:"gear" gorm:"column:gear"`                 // 挡位；
	Mileage          int32     `json:"mileage" gorm:"column:mileage"`           // 累计里程；有效值范围:0~9 999 999(表示0 km~999 999.9 km)最小计量单元;0.1 km“0xFF，0xFF，0xFF,OxFE”表示异常，“0xFF,OxFF,OxFF0xFF”示无效
	Mode             int32     `json:"mode" gorm:"column:mode"`                 // 运行模式；0x01:纯电;0x02;混动;0x03:燃油;0xFE表示异常;0xFF 表示无效
	Resistance       int32     `json:"resistance" gorm:"column:resistance"`     // 绝缘电阻；有效范围0~60 000(表示 0 Q~60 000 k2)最小计量单元;1 k2
	Soc              int32     `json:"soc" gorm:"column:soc"`                   // 有效值范围;0~100(表示 0%~100%),最小计量单元:1%“0xFE”表示异常“0xFF”表示无效
	Speed            int32     `json:"speed" gorm:"column:speed"`               // 车速；有效值范围;0~100(表示 0%~100%),最小计量单元:1%，“0xFE”表示异常“0xFF”表示无效
	Status           int32     `json:"status" gorm:"column:status"`             // 车辆状态；0x01:启动，0x02:熄火，0x03:其他，0xFE表示异常，0xFF表示无效
	Voltage          int32     `json:"voltage" gorm:"column:voltage"`           // 总电压；有效值范围:0~10 000(表示 0 V~1000 V)，小计量单元:0.1V，“0xFF,xFE”示异常，“0xFFxFF”表示无效
	CreatedTime      time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

func (e *VehicleLog) TableName() string {
	return "vehicle_log"
}
