package msgModel

type InfoCateModel struct {
	Vin     string // 当传输车辆数据时，应使用车辆VIN，其字码应符合GB 16735的规定。如传输其他数据，则使用唯一自定义编码
	CmdFlag byte
	MsgDate MsgDate
	Data    []byte
}
