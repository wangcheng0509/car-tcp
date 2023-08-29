package dbmodel

import "time"

// 驱动电机数据Log
type DrivemotorLog struct {
	Id          int64     `json:"id" gorm:"column:id"`
	Vin         string    `json:"vin" gorm:"column:vin"`
	Year        int       `json:"year" gorm:"column:year"`
	Month       int       `json:"month" gorm:"column:month"`
	Day         int       `json:"day" gorm:"column:day"`
	Hour        int       `json:"hour" gorm:"column:hour"`
	Minutes     int       `json:"minutes" gorm:"column:minutes"`
	Seconds     int       `json:"seconds" gorm:"column:seconds"`
	Number      int       `json:"number" gorm:"column:number"`             // 驱动电机个数
	Data        string    `json:"data" gorm:"column:data"`                 // 驱动电机数据
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

func (e *DrivemotorLog) TableName() string {
	return "drivemotor_log"
}
