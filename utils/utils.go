package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go_admin/connect"
	"gopkg.in/gomail.v2"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
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

// 检查文件路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取客户端 IP
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
