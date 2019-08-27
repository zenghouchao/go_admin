package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 生成32位MD5加密值
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
