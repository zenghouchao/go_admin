package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go_admin/connect"
)

// 生成32位MD5加密值
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成API响应JSON数据
func JsonReturn(errCode int32, msg string) []byte {
	res := &connect.Response{
		ErrCode: errCode,
		Msg:     msg,
	}

	jsonString, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error converting json data:\n", err.Error())
	}
	return jsonString

}
