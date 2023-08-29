package api

import (
	"car.open.service/core"
	"car.open.service/ginx"
	"github.com/gin-gonic/gin"
)

// Cmd command to tcp server
type Cmd struct {
}

// CmdReq .
type CmdReq struct {
	Vin string `json:"vin"`
	Cmd string `json:"cmd"`
}

// Send .
func (a *Cmd) Send(c *gin.Context) {
	var cmdReq CmdReq
	err := c.ShouldBind(&cmdReq)
	if err != nil {
		ginx.ResError(c, err, 400)
		return
	}

	err = core.Command(c.Request.Context(), cmdReq.Vin, cmdReq.Cmd)
	if err != nil {
		ginx.ResError(c, err, 500)
		return
	}
	ginx.ResOK(c)
}
