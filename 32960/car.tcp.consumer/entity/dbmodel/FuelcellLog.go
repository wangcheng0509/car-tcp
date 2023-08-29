package dbmodel

import "time"

// 燃料电池数据Log
type FuelcellLog struct {
	Vin               string    `json:"vin" gorm:"column:vin"`
	Year              int32     `json:"year" gorm:"column:year"`
	Month             int32     `json:"month" gorm:"column:month"`
	Day               int32     `json:"day" gorm:"column:day"`
	Hour              int32     `json:"hour" gorm:"column:hour"`
	Minutes           int32     `json:"minutes" gorm:"column:minutes"`
	Seconds           int32     `json:"seconds" gorm:"column:seconds"`
	CellCurrent       int32     `json:"cellCurrent" gorm:"column:cellCurrent"`             // 燃料电池电流
	CellVoltage       int32     `json:"cellVoltage" gorm:"column:cellVoltage"`             // 燃料电池电压
	DcStatus          int32     `json:"dcStatus" gorm:"column:dcStatus"`                   // 高压DC/DC状态
	FuelConsumption   int32     `json:"fuelConsumption" gorm:"column:fuelConsumption"`     // 燃料消耗率
	H_MaxConc         int32     `json:"h_MaxConc" gorm:"column:h_MaxConc"`                 // 氢气最高浓度
	H_ConcSensorCode  int32     `json:"h_ConcSensorCode" gorm:"column:h_ConcSensorCode"`   // 氢气最高浓度传感器代号
	H_MaxPress        int32     `json:"h_MaxPress" gorm:"column:h_MaxPress"`               // 氢气最高压力
	H_PressSensorCode int32     `json:"h_PressSensorCode" gorm:"column:h_PressSensorCode"` // 氢气最高压力传感器代号
	H_MaxTemp         int32     `json:"h_MaxTemp" gorm:"column:h_MaxTemp"`                 // 氢系统中最高温度
	H_TempProbeCode   int32     `json:"h_TempProbeCode" gorm:"column:h_TempProbeCode"`     // 氢系统中最高温度探针代号
	ProbeNum          int32     `json:"probeNum" gorm:"column:probeNum"`                   // 燃料电池温度探针总数
	ProbeTemps        string    `json:"probeTemps" gorm:"column:probeTemps"`               // 探针温度值
	CreatedTime       time.Time `json:"created_time" gorm:"column:created_time"`           // 上报时间
}

func (e *FuelcellLog) TableName() string {
	return "fuelcell_log"
}
