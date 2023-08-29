package model

import (
	"time"
)

// ChargeableTemp 可充电储能装置温度数据
type ChargeableTemp struct {
	Date
	Vin         string    `json:"vin" gorm:"column:vin"`
	Number      int32     `json:"number" gorm:"column:number"`
	Data        string    `json:"data" gorm:"column:data"`                 // 可充电储能子系统电压信息列表
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

// TableName 表名
func (a *ChargeableTemp) TableName() string {
	return "chargeabletemp"
}

// ChargeableTempLog 可充电储能装置温度数据
type ChargeableTempLog struct {
	Id int64 `json:"id" gorm:"column:id"`
	ChargeableTemp
}

// TableName 表名
func (a *ChargeableTempLog) TableName() string {
	return "chargeabletemp_log"
}
