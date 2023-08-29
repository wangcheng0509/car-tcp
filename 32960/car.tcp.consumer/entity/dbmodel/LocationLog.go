package dbmodel

import "time"

// 车辆位置数据Log
type LocationLog struct {
	Vin         string    `json:"vin" gorm:"column:vin"`
	Year        int32     `json:"year" gorm:"column:year"`
	Month       int32     `json:"month" gorm:"column:month"`
	Day         int32     `json:"day" gorm:"column:day"`
	Hour        int32     `json:"hour" gorm:"column:hour"`
	Minutes     int32     `json:"minutes" gorm:"column:minutes"`
	Seconds     int32     `json:"seconds" gorm:"column:seconds"`
	Latitude    int32     `json:"latitude" gorm:"column:latitude"`   // 纬度；以度为单位的纬度值乘以10^6，精确到百万分之一度
	Longitude   int32     `json:"longitude" gorm:"column:longitude"` // 经度；以度为单位的纬度值乘以10^6，精确到百万分之一度
	Status      int32     `json:"status" gorm:"column:status"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

func (e *LocationLog) TableName() string {
	return "location_log"
}
