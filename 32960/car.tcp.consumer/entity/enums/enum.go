package enums

var Encrypt EncryptModel
var Msg MsgModel
var Rsp RspModel

func Init() {
	encryptInit()
	msgInit()
	rspInit()
}

func encryptInit() {
	Encrypt = EncryptModel{
		None: EnumModel{
			Name:  "None",
			Value: 0x01,
		},
		Rsa: EnumModel{
			Name:  "Rsa",
			Value: 0x02,
		},
		Aes: EnumModel{
			Name:  "Aes",
			Value: 0x03,
		},
		Error: EnumModel{
			Name:  "Error",
			Value: 0xFE,
		},
		Invalid: EnumModel{
			Name:  "Invalid",
			Value: 0xFF,
		},
	}
}

func msgInit() {
	Msg = MsgModel{
		VehicleLogin: EnumModel{
			Name:  "VehicleLogin",
			Value: 0x01,
		},
		RealtimeInfo: EnumModel{
			Name:  "RealtimeInfo",
			Value: 0x02,
		},
		ResendInfo: EnumModel{
			Name:  "ResendInfo",
			Value: 0x03,
		},
		VehicleLogout: EnumModel{
			Name:  "VehicleLogout",
			Value: 0xFE,
		},
		Cmd: EnumModel{
			Name:  "Cmd",
			Value: 0x83,
		},
	}
}

func rspInit() {
	Rsp = RspModel{
		Success: EnumModel{
			Name:  "Success",
			Value: 0x01,
		},
		Error: EnumModel{
			Name:  "Error",
			Value: 0x02,
		},
		VINError: EnumModel{
			Name:  "VINError",
			Value: 0x03,
		},
		Cmd: EnumModel{
			Name:  "Cmd",
			Value: 0xFE,
		},
	}
}

type EnumModel struct {
	Name  string
	Value uint8
}

type EncryptModel struct {
	None    EnumModel
	Rsa     EnumModel
	Aes     EnumModel
	Error   EnumModel
	Invalid EnumModel
}

type MsgModel struct {
	VehicleLogin  EnumModel // 车辆登入
	RealtimeInfo  EnumModel // 实时数据
	ResendInfo    EnumModel // 补发数据
	VehicleLogout EnumModel // 车辆登出
	Cmd           EnumModel // 命令下发
}

type RspModel struct {
	Success  EnumModel
	Error    EnumModel
	VINError EnumModel
	Cmd      EnumModel
}
