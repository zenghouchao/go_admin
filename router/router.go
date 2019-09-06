package router

import (
	"go_admin/handler"
	"net/http"
)

func init() {
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
	http.HandleFunc("/article/delete", handler.LoginInterceptor(handler.ArticleDeleteHandler))
	http.HandleFunc("/article/update", handler.LoginInterceptor(handler.ArticleUpdateHandler))
	http.HandleFunc("/article/edit/", handler.LoginInterceptor(handler.ArticleEditPageHandler))

	// 栏目分类
	http.HandleFunc("/cate", handler.LoginInterceptor(handler.CateListHandler))
	http.HandleFunc("/cate/add", handler.LoginInterceptor(handler.CateAddHandler))
	http.HandleFunc("/cate/delete", handler.LoginInterceptor(handler.CateDelHandler))
	http.HandleFunc("/cate/saveState", handler.LoginInterceptor(handler.CateStatusSaveHandler))

	// 邮件配置、
	http.HandleFunc("/email", handler.LoginInterceptor(handler.SetEmailTemplateHandler))
}
