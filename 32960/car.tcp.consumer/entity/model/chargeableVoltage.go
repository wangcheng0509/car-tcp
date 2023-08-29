package model

import (
	"time"
)

// ChargeableVoltage 可充电储能装置电压
type ChargeableVoltage struct {
	Date
	Vin         string    `json:"vin" gorm:"column:vin"`
	Number      int32     `json:"number" gorm:"column:number"`             // 可充电储能子系统个数
	Data        string    `json:"data" gorm:"column:data"`                 // 可充电储能子系统电压信息列表
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

// TableName 表名
func (a *ChargeableVoltage) TableName() string {
	return "chargeablevoltage"
}

// ChargeableVoltageLog 可充电储能装置电压
type ChargeableVoltageLog struct {
	Id int64 `json:"id" gorm:"column:id"`
	ChargeableVoltage
}

// TableName 表名
func (a *ChargeableVoltageLog) TableName() string {
	return "chargeablevoltage_log"
}
