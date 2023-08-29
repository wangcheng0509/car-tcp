package tcpclient

import (
	"encoding/binary"
	"fmt"
)

func getVloginByte() []byte {
	reqParam := Vlogin{
		Year:     23,
		Month:    8,
		Day:      14,
		Hour:     14,
		Minutes:  44,
		Seconds:  30,
		Seq:      30,
		IccId:    "12345678901234567890",
		Num:      4,
		Length:   5,
		EnergyId: "09876543210987654321",
	}
	reqByte := reqParam.Marshal()
	fmt.Println(reqByte)
	return reqByte
}

// 车辆登入
type Vlogin struct {
	Year     int    `json:"year" gorm:"column:year"`
	Month    int    `json:"month" gorm:"column:month"`
	Day      int    `json:"day" gorm:"column:day"`
	Hour     int    `json:"hour" gorm:"column:hour"`
	Minutes  int    `json:"minutes" gorm:"column:minutes"`
	Seconds  int    `json:"seconds" gorm:"column:seconds"`
	Seq      int    `json:"seq" gorm:"column:seq"`           // 车辆登录的流水号
	IccId    string `json:"iccId" gorm:"column:iccId"`       // iccid
	Num      int    `json:"num" gorm:"column:num"`           // 可充电储能子系统数
	Length   int    `json:"length" gorm:"column:length"`     // 可充电储能系统编码长度。
	EnergyId string `json:"energyId" gorm:"column:energyId"` // 可充电储能系统编码列表。
}

func (v Vlogin) Marshal() []byte {
	var bytes []byte
	bytes = append(bytes, byte(v.Year))
	bytes = append(bytes, byte(v.Month))
	bytes = append(bytes, byte(v.Day))
	bytes = append(bytes, byte(v.Hour))
	bytes = append(bytes, byte(v.Minutes))
	bytes = append(bytes, byte(v.Seconds))

	binary.BigEndian.PutUint16(bytes[0:2], uint16(v.Seq))
	bytes = append(bytes, []byte(v.IccId)...)
	bytes = append(bytes, byte(v.Num))
	bytes = append(bytes, byte(v.Length))
	bytes = append(bytes, []byte(v.EnergyId)...)
	return bytes
}
