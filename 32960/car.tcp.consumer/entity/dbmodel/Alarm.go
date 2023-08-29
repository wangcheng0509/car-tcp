package dbmodel

import "time"

// 报警数据
type Alarm struct {
	Vin                       string    `json:"vin" gorm:"column:vin"`
	Year                      int       `json:"year" gorm:"column:year"`
	Month                     int       `json:"month" gorm:"column:month"`
	Day                       int       `json:"day" gorm:"column:day"`
	Hour                      int       `json:"hour" gorm:"column:hour"`
	Minutes                   int       `json:"minutes" gorm:"column:minutes"`
	Seconds                   int       `json:"seconds" gorm:"column:seconds"`
	FaultChargeableDeviceNum  int       `json:"faultChargeableDeviceNum" gorm:"column:faultChargeableDeviceNum"`   // 可充电储能装置故障总数
	FaultChargeableDeviceList string    `json:"faultChargeableDeviceList" gorm:"column:faultChargeableDeviceList"` // 可充电储能装置故障代码列表
	FaultDriveMotorNum        int       `json:"faultDriveMotorNum" gorm:"column:faultDriveMotorNum"`               // 驱动电机故障总数
	FaultDriveMotorList       string    `json:"faultDriveMotorList" gorm:"column:faultDriveMotorList"`             // 驱动电机故障代码列表
	FaultEngineNum            int       `json:"faultEngineNum" gorm:"column:faultEngineNum"`                       // 发动机故障总数
	FaultEngineList           string    `json:"faultEngineList" gorm:"column:faultEngineList"`                     // 发动机故障列表
	FaultOthersNum            int       `json:"faultOthersNum" gorm:"column:faultOthersNum"`                       // 其他故障总数
	FaultOthersList           string    `json:"faultOthersList" gorm:"column:faultOthersList"`                     // 其他故障代码列表
	GeneralAlarmFlag          string    `json:"generalAlarmFlag" gorm:"column:generalAlarmFlag"`                   // 通用报警标志
	MaxAlarmLevel             int       `json:"maxAlarmLevel" gorm:"column:maxAlarmLevel"`                         // 最高报警等级
	CreatedTime               time.Time `json:"created_time" gorm:"column:created_time"`                           // 上报时间
}

func (e *Alarm) TableName() string {
	return "alarm"
}
