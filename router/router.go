package router

import (
	"github.com/dchest/captcha"
	"github.com/go_admin/handler"
	"net/http"
)

func init() {
	// Static file routing registration
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler.IndexHandler)
	http.Handle("/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))
	http.HandleFunc("/login", handler.DoLogin)
	http.HandleFunc("/upload", handler.LoginInterceptor(handler.UploadFile))
	http.HandleFunc("/home", handler.LoginInterceptor(handler.AdminHandler))
	http.HandleFunc("/home/welcome", handler.LoginInterceptor(handler.WelcomeHandler))
	http.HandleFunc("/logout", handler.LoginInterceptor(handler.LogoutHandler))
	http.HandleFunc("/adminUser/changePass", handler.LoginInterceptor(handler.AdminPassChange))

	// Article Manage
	http.HandleFunc("/article", handler.LoginInterceptor(handler.ArticleListHandler))
	http.HandleFunc("/article/add", handler.LoginInterceptor(handler.ArticleAddHandler))
	http.HandleFunc("/article/delete", handler.LoginInterceptor(handler.ArticleDeleteHandler))
	http.HandleFunc("/article/update", handler.LoginInterceptor(handler.ArticleUpdateHandler))
	http.HandleFunc("/article/edit/", handler.LoginInterceptor(handler.ArticleEditPageHandler))

	// Article Category
	http.HandleFunc("/cate", handler.LoginInterceptor(handler.CateListHandler))
	http.HandleFunc("/cate/add", handler.LoginInterceptor(handler.CateAddHandler))
	http.HandleFunc("/cate/delete", handler.LoginInterceptor(handler.CateDelHandler))
	http.HandleFunc("/cate/saveState", handler.LoginInterceptor(handler.CateStatusSaveHandler))

	// Email Send
	http.HandleFunc("/email", handler.LoginInterceptor(handler.SetEmailTemplateHandler))
	http.HandleFunc("/email/send", handler.LoginInterceptor(handler.SendEmailHandeler))

	// IM chat
	http.HandleFunc("/imChat", handler.LoginInterceptor(handler.ImChatHandler))
}
