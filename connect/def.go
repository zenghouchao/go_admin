package connect

// 后台用户加密盐
var Salt = "fhsdnfgdjrweirwe1324186asdasdqw"

const (
	OK_API     = 0
	ERR_API    = 1
	EMAIL_HOST = ""
	EMAIL_PORT = ""
	EMAIL_USER = ""
	EMAIL_PASS = ""
)

var PageSize = 10

// API response return
type Response struct {
	ErrCode int32                    `json:"errCode"`
	Msg     string                   `json:"msg"`
	Data    []map[string]interface{} `json:"data,omitempty"`
}

// 文章分类表结构
type Cate struct {
	Id     string
	Name   string
	Status string
}

// 文章表结构
type Article struct {
	Id      string
	Cate_id string
	Title   string
	Content string
	Time    int64
	Status  string
	Author  string
	Pubdate string // Time字段的时间戳格式化
}

// 用户表结构体
type User struct {
	Id       int
	Name  string
	Pass  string
	Time  int
}
