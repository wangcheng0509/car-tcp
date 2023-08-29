package dbmodel

import "time"

// 车辆登入Log
type VloginLog struct {
	Vin         string    `json:"vin" gorm:"column:vin"`
	Year        int32     `json:"year" gorm:"column:year"`
	Month       int32     `json:"month" gorm:"column:month"`
	Day         int32     `json:"day" gorm:"column:day"`
	Hour        int32     `json:"hour" gorm:"column:hour"`
	Minutes     int32     `json:"minutes" gorm:"column:minutes"`
	Seconds     int32     `json:"seconds" gorm:"column:seconds"`
	Seq         int32     `json:"seq" gorm:"column:seq"`                   // 车辆登录的流水号
	IccId       string    `json:"iccId" gorm:"column:iccId"`               // iccid
	Num         int32     `json:"num" gorm:"column:num"`                   // 可充电储能子系统数
	Length      int32     `json:"length" gorm:"column:length"`             // 可充电储能系统编码长度
	EnergyId    string    `json:"energyId" gorm:"column:energyId"`         // 可充电储能系统编码列表
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

func (e *VloginLog) TableName() string {
	return "vlogin_log"
}
