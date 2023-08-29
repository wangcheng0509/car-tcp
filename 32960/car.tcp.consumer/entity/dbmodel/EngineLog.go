package dbmodel

import "time"

// 发动机数据Log
type EngineLog struct {
	Id              int64     `json:"id" gorm:"column:id"`
	Vin             string    `json:"vin" gorm:"column:vin"`
	Year            int       `json:"year" gorm:"column:year"`
	Month           int       `json:"month" gorm:"column:month"`
	Day             int       `json:"day" gorm:"column:day"`
	Hour            int       `json:"hour" gorm:"column:hour"`
	Minutes         int       `json:"minutes" gorm:"column:minutes"`
	Seconds         int       `json:"seconds" gorm:"column:seconds"`
	CrankshaftSpeed int       `json:"crankshaftSpeed" gorm:"column:crankshaftSpeed"` // 曲轴转速；有效范围：0-60 000（表示0 r/min-60 000 r/min)，最小计量单元：1 r/min,0xFF/0xFE表示异常，0xFF,0xFF表示无效
	FuelConsumption float32   `json:"fuelConsumption" gorm:"column:fuelConsumption"` // 燃料消耗率；有效范围:0~60 000(表示0r/min一60 000 r/min)，最小计量单元:1r/min，0xFF/OxFE表示异常，0xFF表示无效
	Status          int       `json:"status" gorm:"column:status"`                   // 发动机状态；0x01启动状态；0x02关闭状态，0xFE表示异常，0xFF表示无效
	CreatedTime     time.Time `json:"created_time" gorm:"column:created_time"`       // 上报时间
}

func (e *EngineLog) TableName() string {
	return "engine_log"
}
