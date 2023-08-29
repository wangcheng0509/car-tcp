package tcpclient

func getVlogoutBytes() (data []byte) {
	reqParam := dateModel{
		Year:    23,
		Month:   8,
		Day:     16,
		Hour:    17,
		Minutes: 17,
		Seconds: 30,
	}
	data = reqParam.Marshal()
	data = append(data, []byte("30")...)
	return
}
