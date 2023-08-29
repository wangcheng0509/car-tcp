package handleEntity

import (
	"encoding/json"
	"time"
)

type CacheModel struct {
	Id          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Vin         string    `json:"vin" gorm:"column:vin"`
	Host        string    `json:"host" gorm:"column:host"`
	Port        string    `json:"port" gorm:"column:port"`
	ConnStr     int       `json:"connStr" gorm:"column:connStr"`
	Status      int       `json:"status" gorm:"column:status"`
	OnlineDate  time.Time `json:"onlineDate" gorm:"column:onlineDate"`
	OfflineDate time.Time `json:"offlineDate" gorm:"column:offlineDate"`
}

func (e *CacheModel) TableName() string {
	return "vehicle_conn"
}
func (g *CacheModel) MarshalBinary() (data []byte, err error) {
	return json.Marshal(g)
}
