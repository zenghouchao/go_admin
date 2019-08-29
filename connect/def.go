package connect

// 后台用户加密盐
var Salt = "fhsdnfgdjrweirwe1324186asdasdqw"

var PageSize = 20

// API响应返回值
type Response struct {
	ErrCode int32       `json:"errCode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 文章分类表结构
type Cate struct {
	Id     string
	Name   string
	Status string
}
