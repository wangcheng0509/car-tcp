package handleEntity

type VloginModel struct {
	Vin                string // Vin
	Host               string // 链接所在服务器
	ConnStr            string // 链接名称
	Year               uint8  // 消息年
	Month              uint8  // 消息月
	Day                uint8  // 消息日
	Hour               uint8  // 消息时
	Minutes            uint8  // 消息分
	Seconds            uint8  // 消息秒
	LoginNum           uint16 // 车辆登录的流水号。
	IccId              string // IccId
	SubSystemNumber    uint8  // 可充电储能子系统数。
	SystemEncodeLength uint8  // 可充电储能系统编码长度。
	SystemEncode       string // 可充电储能系统编码列表。
}
