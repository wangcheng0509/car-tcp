package model

import "time"

// 车辆链接注册表
type VehicleConn struct {
	Id          int       `json:"id" gorm:"column:id"`
	Vin         string    `json:"vin" gorm:"column:vin"`
	Host        string    `json:"host" gorm:"column:host"`
	Port        string    `json:"port" gorm:"column:port"`
	ConnStr     string    `json:"connStr" gorm:"column:connStr"`
	Status      int       `json:"status" gorm:"column:status"` // 1在线
	OnlineDate  time.Time `json:"onlineDate" gorm:"column:onlineDate"`
	OfflineDate time.Time `json:"offlineDate" gorm:"column:offlineDate"`
}

func (e *VehicleConn) TableName() string {
	return "vehicle_conn"
}
