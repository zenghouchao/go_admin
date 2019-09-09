package handler

import (
	"fmt"
	"github.com/astaxie/beego/session"
	"github.com/dchest/captcha"
	"go_admin/connect"
	"go_admin/dao"
	"go_admin/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	globalSessions *session.Manager
)

func init() {
	sc := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
	}
	globalSessions, _ = session.NewManager("memory", sc)
	go globalSessions.GC()
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles("./template/login.html")

	if err != nil {
		panic("加载模板失败")
	}
	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, nil)
}

func DoLogin(w http.ResponseWriter, r *http.Request) {
	var response []byte
	if r.Method == "POST" {
		r.ParseForm()
		user := r.Form.Get("username")
		pass := r.Form.Get("password")
		captchaCode := r.Form.Get("captcha")
		// 验证码
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)

		captchaId := sess.Get("captcha")
		fmt.Println(captchaId)

		ok := captcha.VerifyString("12", captchaCode)
		if !ok {
			response = utils.JsonReturn(connect.ERR_API, "验证码错误！")
			w.Write(response)
			return
		}

		pass += connect.Salt
		checkd := dao.AdminLogin(user, utils.Md5(pass))
		if !checkd {
			//io.WriteString(w, "用户名或密码错误")
			response = utils.JsonReturn(connect.ERR_API, "用户名或密码错误！")
			w.Write(response)
			return
		}
		// 登陆成功
		err := sess.Set("userInfo", user)
		if err != nil {
			panic(err.Error())
		}
		http.Redirect(w, r, "/home", http.StatusFound)

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/home.html")
	if err != nil {
		panic("加载后台模板失败")
	}
	w.Header().Set("Content-Type", "text/html")
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	adminInfo := sess.Get("userInfo")

	if err := tpl.Execute(w, adminInfo); err != nil {
		fmt.Println("Template rendering failed !")
	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/welcome.html")
	if err != nil {
		panic("加载后台模板失败")
	}
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	adminInfo := sess.Get("userInfo")

	w.Header().Set("Content-Type", "text/html")
	params := map[string]interface{}{
		"nowTime":   time.Now().Format("2006-01-02 15:04:05"),
		"adminInfo": adminInfo,
	}
	tpl.Execute(w, params)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// 清除SESSION
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	err := sess.Delete("userInfo")
	if err != nil {
		fmt.Println("session delete error: " + err.Error())
		return
	}

	//http.Redirect(w, r, "/", http.StatusFound) 会调用auth验证此处多余
	return
}

func AdminPassChange(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var response []byte
		if err := dao.AdminPassChange(r); err != nil {
			response = utils.JsonReturn(connect.ERR_API, err.Error())
		} else {
			response = utils.JsonReturn(connect.OK_API, "修改管理员密码成功")
			defer LogoutHandler(w, r)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(response)))
		w.Write(response)
		return

	} else {
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)
		adminInfo := sess.Get("userInfo")
		tpl, err := template.ParseFiles("./template/changePass.html")
		if err != nil {
			panic("loading change admin user template error~")
			return
		}
		tpl.Execute(w, adminInfo)
	}
}

// 验证码
func CaptchaHandler(w http.ResponseWriter, r *http.Request) {
	id := captcha.New()
	err := captcha.WriteImage(w, id, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		fmt.Println("生成验证码错误：", err.Error())
		return
	}
	// 给ID存储在SESSION中
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	if err = sess.Set("captcha", id); err != nil {
		log.Println("captcha session store error")
		return
	}
}
