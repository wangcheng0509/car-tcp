package tools

import (
	"github.com/aceld/zinx/ziface"
)

// 验证校验码
func CheckCode(request ziface.IRequest) {
	data := request.GetData()
	if len(data) < 25 {
		request.Abort()
	}
	_code := data[2]
	for _, v := range data[3 : len(data)-2] {
		_code = _code ^ v
	}
	if _code != data[len(data)-1] {
		data[3] = 0x02
		request.GetConnection().SendMsg(uint8(data[2]), data)
		request.Abort()
	}
}

// 验证校验码
func GetCode(data []byte) byte {
	if len(data) < 25 {
		return 0
	}
	_code := data[2]
	for _, v := range data[3 : len(data)-1] {
		_code = _code ^ v
	}
	return _code
}
