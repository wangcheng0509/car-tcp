package dbmodel

import "time"

// 极值数据
type Extreme struct {
	Vin                       string    `json:"vin" gorm:"column:vin"`
	Year                      int       `json:"year" gorm:"column:year"`
	Month                     int       `json:"month" gorm:"column:month"`
	Day                       int       `json:"day" gorm:"column:day"`
	Hour                      int       `json:"hour" gorm:"column:hour"`
	Minutes                   int       `json:"minutes" gorm:"column:minutes"`
	Seconds                   int       `json:"seconds" gorm:"column:seconds"`
	MaxBatteryVoltage         int       `json:"maxBatteryVoltage" gorm:"column:maxBatteryVoltage"`                 // 电池单体电压最高值
	MaxTemp                   int       `json:"maxTemp" gorm:"column:maxTemp"`                                     // 最高温度值
	MaxTempProbeNo            int       `json:"maxTempProbeNo" gorm:"column:maxTempProbeNo"`                       // 最高温度探针序号
	MaxTempSubsysNo           int       `json:"maxTempSubsysNo" gorm:"column:maxTempSubsysNo"`                     // 最高温度子系统号
	MaxVoltageBatteryCode     int       `json:"maxVoltageBatteryCode" gorm:"column:maxVoltageBatteryCode"`         // 最高电压电池单体代号
	MaxVoltageBatterySubsysNo int       `json:"maxVoltageBatterySubsysNo" gorm:"column:maxVoltageBatterySubsysNo"` // 最高电压电池子系统号
	MinBatteryVoltage         int       `json:"minBatteryVoltage" gorm:"column:minBatteryVoltage"`                 // 电池单体电压最低值
	MinTemp                   int       `json:"minTemp" gorm:"column:minTemp"`                                     // 最低温度值
	MinTempProbeNo            int       `json:"minTempProbeNo" gorm:"column:minTempProbeNo"`                       // 最低温度探针序号
	MinTempSubsysNo           int       `json:"minTempSubsysNo" gorm:"column:minTempSubsysNo"`                     // 最低温度子系统号
	MinVoltageBatteryCode     int       `json:"minVoltageBatteryCode" gorm:"column:minVoltageBatteryCode"`         // 最低电压电池单体代号
	MinVoltageBatterySubsysNo int       `json:"minVoltageBatterySubsysNo" gorm:"column:minVoltageBatterySubsysNo"` // 最低电压电池子系统号
	CreatedTime               time.Time `json:"created_time" gorm:"column:created_time"`                           // 上报时间
}

func (e *Extreme) TableName() string {
	return "extreme"
}
