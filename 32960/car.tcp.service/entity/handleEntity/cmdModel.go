package handleEntity

type CmdModel struct {
	ConnId uint64 `json:"connId"`
	Msg    []byte `json:"msg"`
}
