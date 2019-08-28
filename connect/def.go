package connect

// 后台用户加密盐
var Salt = "fhsdnfgdjrweirwe1324186asdasdqw"

// API响应返回值
type Response struct {
	ErrCode int32       `json:"errCode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
