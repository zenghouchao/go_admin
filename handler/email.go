package handler

import (
	"fmt"
	"html/template"
	"net/http"
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

}
