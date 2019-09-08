package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go_admin/connect"
	"gopkg.in/gomail.v2"
	"strconv"
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

// 发送邮件
func SendMail(mailFrom string, mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": connect.EMAIL_USER,
		"pass": connect.EMAIL_PASS,
		"host": connect.EMAIL_HOST,
		"port": connect.EMAIL_PORT,
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", mailFrom)
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}
