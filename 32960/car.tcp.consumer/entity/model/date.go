package model

// Date 日期
type Date struct {
	Year    int32 `json:"year" gorm:"column:year"`
	Month   int32 `json:"month" gorm:"column:month"`
	Day     int32 `json:"day" gorm:"column:day"`
	Hour    int32 `json:"hour" gorm:"column:hour"`
	Minutes int32 `json:"minutes" gorm:"column:minutes"`
	Seconds int32 `json:"seconds" gorm:"column:seconds"`
}
