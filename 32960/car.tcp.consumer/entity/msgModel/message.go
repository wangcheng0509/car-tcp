package msgModel

import (
	"bytes"
)

type Message struct {
	Head string // 固定为ASCII字符‘# #'，用“0x23, 0x23"表示
	// 0x01 车辆登入 上行
	// 0x02 实时信息上报 上行
	// 0x03 补发信息上报 上行
	// 0x04 车辆登出 上行
	// 0x05 平台登入 上行
	// 0x06 平台登出 上行
	// 0x07~0x08 终端数据预留 上行
	// 0x09~0x7F 上行数据系统预留 上行
	// 0x80~0x82 终端数据预留 下行
	// 0x83~0xBF 下行数据系统预留 下行
	// oxC0~0xFE 平台交换自定义数据 自定义
	CmdCFlag uint8  // 命令标识
	CmdRsp   uint8  // 应答标志
	Vin      string // 当传输车辆数据时，应使用车辆VIN，其字码应符合GB 16735的规定。如传输其他数据，则使用唯一自定义编码
	Encrypt  uint8  // 0x01 数据不加密,0x02 RSA算法加密,0x03 AES128加密,0xFE表示异常,0xFF表示无效,其他预留
	Len      uint16 // 数据单元长度，双字节整型，范围0~65531
	Data     []byte // 数据单元
	Code     uint8  // 采用 BCC(异或校验)法，校验范围从命令单元的第一个字节开始，同后一字节异或，直到校验码前一字节为止，校验码占用一个字节，当数据单元存在加密时，应先加密后校验，先校验后解密
}

func (m Message) Marshal() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(m.Head))

	var cmdbyte []byte
	cmdbyte = append(cmdbyte, byte(m.CmdCFlag))
	cmdbyte = append(cmdbyte, byte(m.CmdRsp))
	buffer.Write(cmdbyte)

	buffer.Write([]byte(m.Vin))

	var encryptbyte []byte
	encryptbyte = append(cmdbyte, byte(m.Encrypt))
	encryptbyte[1] = byte(m.Len >> 8)
	encryptbyte[2] = byte(m.Len)
	buffer.Write(encryptbyte)

	buffer.Write([]byte(m.Data))
	return buffer.Bytes()
}
