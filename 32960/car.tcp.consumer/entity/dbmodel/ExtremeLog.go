package dbmodel

import "time"

// 极值数据Log
type ExtremeLog struct {
	Vin                       string    `json:"vin" gorm:"column:vin"`
	Year                      int32     `json:"year" gorm:"column:year"`
	Month                     int32     `json:"month" gorm:"column:month"`
	Day                       int32     `json:"day" gorm:"column:day"`
	Hour                      int32     `json:"hour" gorm:"column:hour"`
	Minutes                   int32     `json:"minutes" gorm:"column:minutes"`
	Seconds                   int32     `json:"seconds" gorm:"column:seconds"`
	MaxBatteryVoltage         int32     `json:"maxBatteryVoltage" gorm:"column:maxBatteryVoltage"`                 // 电池单体电压最高值
	MaxTemp                   int32     `json:"maxTemp" gorm:"column:maxTemp"`                                     // 最高温度值
	MaxTempProbeNo            int32     `json:"maxTempProbeNo" gorm:"column:maxTempProbeNo"`                       // 最高温度探针序号
	MaxTempSubsysNo           int32     `json:"maxTempSubsysNo" gorm:"column:maxTempSubsysNo"`                     // 最高温度子系统号
	MaxVoltageBatteryCode     int32     `json:"maxVoltageBatteryCode" gorm:"column:maxVoltageBatteryCode"`         // 最高电压电池单体代号
	MaxVoltageBatterySubsysNo int32     `json:"maxVoltageBatterySubsysNo" gorm:"column:maxVoltageBatterySubsysNo"` // 最高电压电池子系统号
	MinBatteryVoltage         int32     `json:"minBatteryVoltage" gorm:"column:minBatteryVoltage"`                 // 电池单体电压最低值
	MinTemp                   int32     `json:"minTemp" gorm:"column:minTemp"`                                     // 最低温度值
	MinTempProbeNo            int32     `json:"minTempProbeNo" gorm:"column:minTempProbeNo"`                       // 最低温度探针序号
	MinTempSubsysNo           int32     `json:"minTempSubsysNo" gorm:"column:minTempSubsysNo"`                     // 最低温度子系统号
	MinVoltageBatteryCode     int32     `json:"minVoltageBatteryCode" gorm:"column:minVoltageBatteryCode"`         // 最低电压电池单体代号
	MinVoltageBatterySubsysNo int32     `json:"minVoltageBatterySubsysNo" gorm:"column:minVoltageBatterySubsysNo"` // 最低电压电池子系统号
	CreatedTime               time.Time `json:"created_time" gorm:"column:created_time"`                           // 上报时间
}

func (e *ExtremeLog) TableName() string {
	return "extreme_log"
}
