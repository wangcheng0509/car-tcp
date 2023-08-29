package tools

import (
	"bytes"
	"encoding/binary"
)

type Msg struct {
	Head   byte
	Length int32
	Type   int16
	Data   string
	Tail   byte
}

func Encode(message Msg) ([]byte, error) {
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.BigEndian, message.Head)
	if err != nil {
		return nil, err
	}
	// 写入长度
	err = binary.Write(pkg, binary.BigEndian, message.Length)
	if err != nil {
		return nil, err
	}
	// 写入类型
	err = binary.Write(pkg, binary.BigEndian, message.Type)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.BigEndian, []byte(message.Data))
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.BigEndian, message.Tail)
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
