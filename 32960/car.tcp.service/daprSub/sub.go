package daprsub

import (
	"encoding/json"
	"io"
	"net/http"

	"car.tcp.service/entity/handleEntity"
	"car.tcp.service/handle/service"
)

func Init() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/cmdDown", eventHandler)
	http.ListenAndServe(":8080", myMux)
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var res resMsg
		var req handleEntity.CmdModel
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &req)
		if err == nil {
			if code, err := service.CmdDown(req); code == 2 {
				res.Code = "500"
				res.Data = err.Error()
			} else if code == 1 {
				res.Code = "500"
				res.Data = err.Error()
			} else {
				res.Code = "200"
				res.Data = "success"
			}
			resstr, _ := json.Marshal(res)
			w.Write(resstr)
			return
		}
		res.Code = "500"
		res.Data = err.Error()
		resstr, _ := json.Marshal(res)
		w.Write(resstr)
	}
}

type resMsg struct {
	Code string      `json:"code"` // 响应代码
	Data interface{} `json:"data"` // 响应数据
}
