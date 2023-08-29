package model

import (
	"time"
)

// Engine 发动机数据
type Engine struct {
	Date

	Vin             string    `json:"vin" gorm:"column:vin"`
	CrankshaftSpeed int32     `json:"crankshaftSpeed" gorm:"column:crankshaftSpeed"` // 曲轴转速；有效范围：0-60 000（表示0 r/min-60 000 r/min)，最小计量单元：1 r/min,0xFF/0xFE表示异常，0xFF,0xFF表示无效
	FuelConsumption int32     `json:"fuelConsumption" gorm:"column:fuelConsumption"` // 燃料消耗率；有效范围:0~60 000(表示0r/min一60 000 r/min)，最小计量单元:1r/min，0xFF/OxFE表示异常，0xFF表示无效
	Status          int32     `json:"status" gorm:"column:status"`                   // 发动机状态；0x01启动状态；0x02关闭状态，0xFE表示异常，0xFF表示无效
	CreatedTime     time.Time `json:"created_time" gorm:"column:created_time"`       // 上报时间
}

// TableName 表名
func (e *Engine) TableName() string {
	return "engine"
}

// 发动机数据Log
type EngineLog struct {
	Id int32 `json:"id" gorm:"column:id"`
	Engine
}

func (e *EngineLog) TableName() string {
	return "engine_log"
}
