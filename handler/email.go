package handler

import (
	"fmt"
	"go_admin/connect"
	"go_admin/utils"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func SetEmailTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/email/set.html")
	if err != nil {
		fmt.Println("Loading email template error:", err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, nil)
}

func SendEmailHandeler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postForm := r.PostForm
	fmt.Println(postForm)
	m := gomail.NewMessage()
	m.SetHeader("From", postForm.Get("fromUser"))
	m.SetHeader("To", postForm.Get("toUser"))
	m.SetHeader("Subject", postForm.Get("subject"))
	m.SetBody("text/html", postForm.Get("content"))

	d := gomail.NewDialer(connect.EMAIL_HOST, connect.EMAIL_PORT, connect.EMAIL_USER, connect.EMAIL_PASS)

	var response []byte
	if err := d.DialAndSend(m); err != nil {
		log.Println("send email error:", err.Error())
		response = utils.JsonReturn(connect.ERR_API, err.Error())
	} else {
		response = utils.JsonReturn(connect.OK_API, "发送邮件成功")
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	w.Write(response)
}
