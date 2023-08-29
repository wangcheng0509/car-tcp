package dbmodel

import "time"

// 车辆登入
type Vlogin struct {
	Vin         string    `json:"vin" gorm:"column:vin"`
	Year        int       `json:"year" gorm:"column:year"`
	Month       int       `json:"month" gorm:"column:month"`
	Day         int       `json:"day" gorm:"column:day"`
	Hour        int       `json:"hour" gorm:"column:hour"`
	Minutes     int       `json:"minutes" gorm:"column:minutes"`
	Seconds     int       `json:"seconds" gorm:"column:seconds"`
	Seq         int       `json:"seq" gorm:"column:seq"`                   // 车辆登录的流水号
	IccId       string    `json:"iccId" gorm:"column:iccId"`               // iccid
	Num         int       `json:"num" gorm:"column:num"`                   // 可充电储能子系统数
	Length      int       `json:"length" gorm:"column:length"`             // 可充电储能系统编码长度。
	EnergyId    string    `json:"energyId" gorm:"column:energyId"`         // 可充电储能系统编码列表。
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

func (e *Vlogin) TableName() string {
	return "vlogin"
}
