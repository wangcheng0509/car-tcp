package dbmodel

import "time"

// 车辆位置数据
type Location struct {
	Vin         string    `json:"vin" gorm:"column:vin"`
	Year        int       `json:"year" gorm:"column:year"`
	Month       int       `json:"month" gorm:"column:month"`
	Day         int       `json:"day" gorm:"column:day"`
	Hour        int       `json:"hour" gorm:"column:hour"`
	Minutes     int       `json:"minutes" gorm:"column:minutes"`
	Seconds     int       `json:"seconds" gorm:"column:seconds"`
	Latitude    int       `json:"latitude" gorm:"column:latitude"`   // 纬度；以度为单位的纬度值乘以10^6，精确到百万分之一度
	Longitude   int       `json:"longitude" gorm:"column:longitude"` // 经度；以度为单位的纬度值乘以10^6，精确到百万分之一度
	Status      int       `json:"status" gorm:"column:status"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

func (e *Location) TableName() string {
	return "location"
}
