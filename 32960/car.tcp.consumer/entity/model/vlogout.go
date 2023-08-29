package model

import (
	"time"
)

// Vlogout 车辆登出
type Vlogout struct {
	Vin string `json:"vin" gorm:"column:vin"`
	Seq int    `json:"seq" gorm:"column:seq"` // 车辆登录的流水号
	Date
	CreatedTime time.Time `json:"createdTime" gorm:"column:createdTime"`
}

// TableName 表名
func (a *Vlogout) TableName() string {
	return "vlogout"
}

type VlogoutLog struct {
	Id int32 `json:"id" gorm:"column:id"` // id
	Vlogout
}

// TableName 表名
func (a *VlogoutLog) TableName() string {
	return "vlogout_log"
}
