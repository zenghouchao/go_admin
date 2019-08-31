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
	// 文章管理
	http.HandleFunc("/article", handler.LoginInterceptor(handler.ArticleListHandler))
	http.HandleFunc("/article/add", handler.LoginInterceptor(handler.ArticleAddHandler))

	// 栏目分类
	http.HandleFunc("/cate", handler.LoginInterceptor(handler.CateListHandler))
	http.HandleFunc("/cate/add", handler.LoginInterceptor(handler.CateAddHandler))
	http.HandleFunc("/cate/delete", handler.LoginInterceptor(handler.CateDelHandler))
	http.HandleFunc("/cate/saveState", handler.LoginInterceptor(handler.CateStatusSaveHandler))

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("The HTTP server failed to start:\n", err.Error())
	}
}
