package model

import (
	"time"
)

// 驱动电机数据
type Drivemotors struct {
	Vin    string `json:"vin" gorm:"column:vin"`
	Number int32  `json:"number" gorm:"column:number"` // 驱动电机数量
	Data   string `json:"data" gorm:"column:data"`     // 驱动电机列表
	Date
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 上报时间
}

type Drivemotor struct {
	No           int32 `json:"no" gorm:"column:no"`                     // 驱动电机序号
	CtrlTemp     int32 `json:"ctrlTemp" gorm:"column:ctrlTemp"`         // 驱动电机控制器温度
	DcBusCurrent int32 `json:"dcBusCurrent" gorm:"column:dcBusCurrent"` // 电机控制器直流母线电流
	InputVoltage int32 `json:"inputVoltage" gorm:"column:inputVoltage"` // 电机控制器输人电压
	MotorTemp    int32 `json:"motorTemp" gorm:"column:motorTemp"`       // 驱动电机温度
	Rotating     int32 `json:"rotating" gorm:"column:rotating"`         // 驱动电机转速
	Status       int32 `json:"status" gorm:"column:status"`             // 驱动电机状态
	Torque       int32 `json:"torque" gorm:"column:torque"`             // 驱动电机转矩
}

func (e *Drivemotors) TableName() string {
	return "drivemotor"
}

type DrivemotorsLog struct {
	Id int32 `json:"id" gorm:"column:id"`
	Drivemotors
}

func (e *DrivemotorsLog) TableName() string {
	return "drivemotor_log"
}
