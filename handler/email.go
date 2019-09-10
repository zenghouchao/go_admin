package handler

import (
	"fmt"
	"github.com/go_admin/connect"
	"github.com/go_admin/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	if r.Method == "POST" {
		r.ParseForm()
		postForm := r.PostForm
		subject := strings.TrimSpace(postForm.Get("subject"))
		body := strings.TrimSpace(postForm.Get("content"))
		// 可能有多个收件人
		mailTo := strings.Split(postForm.Get("toUser"), ";")
		mailFrom := strings.TrimSpace(postForm.Get("fromUser"))

		err := utils.SendMail(mailFrom, mailTo, subject, body)
		//fmt.Println("send email error:", err)

		var response []byte
		if err != nil {
			log.Println("send email error:", err.Error())
			response = utils.JsonReturn(connect.ERR_API, err.Error())
		} else {
			response = utils.JsonReturn(connect.OK_API, "发送邮件成功")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(response)))
		w.Write(response)
	}
}
