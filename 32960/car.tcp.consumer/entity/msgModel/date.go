package msgModel

type MsgDate struct {
	Year    int `json:"year" gorm:"column:year"`
	Month   int `json:"month" gorm:"column:month"`
	Day     int `json:"day" gorm:"column:day"`
	Hour    int `json:"hour" gorm:"column:hour"`
	Minutes int `json:"minutes" gorm:"column:minutes"`
	Seconds int `json:"seconds" gorm:"column:seconds"`
}
