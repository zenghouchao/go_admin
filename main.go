package main

import (
	"fmt"
	"go_admin/handler"
	"net/http"
)

func main() {
	// 静态文件路由注册
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", handler.DoLogin)
	http.HandleFunc("/home", handler.LoginInterceptor(handler.AdminHandler))
	http.HandleFunc("/home/welcome", handler.LoginInterceptor(handler.WelcomeHandler))
	http.HandleFunc("/logout", handler.LoginInterceptor(handler.LogoutHandler))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP服务器启动失败", err.Error())
	}
}
