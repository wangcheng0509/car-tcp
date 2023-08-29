package tools

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func GetUInt16(data []byte) int {
	if len(data) != 2 {
		return 0
	}
	in := bytes.NewBuffer(data)
	var value uint16
	binary.Read(in, binary.BigEndian, &value)
	return int(value)
}

func GetUInt32(data []byte) int {
	if len(data) != 4 {
		return 0
	}
	in := bytes.NewBuffer(data)
	var value uint32
	binary.Read(in, binary.BigEndian, &value)
	return int(value)
}

func GetByteFromUint16(data int) ([]byte, error) {
	var result []byte
	var bytes [2]byte
	if data > 65535 {
		return result, errors.New(">65535")
	}
	binary.BigEndian.PutUint16(bytes[0:2], uint16(data))
	result = append(result, bytes[0])
	result = append(result, bytes[1])
	return result, nil
}

func GetByteFromUint32(data int) ([]byte, error) {
	var result []byte
	var bytes [4]byte
	if data > 4294967295 {
		return result, errors.New(">4294967295")
	}
	binary.BigEndian.PutUint32(bytes[0:4], uint32(data))
	result = append(result, bytes[0])
	result = append(result, bytes[1])
	result = append(result, bytes[2])
	result = append(result, bytes[3])
	return result, nil
}
