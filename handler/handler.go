package handler

import (
	"fmt"
		"github.com/astaxie/beego/session"
	"github.com/dchest/captcha"
	"github.com/go_admin/connect"
	"github.com/go_admin/dao"
	"github.com/go_admin/utils"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

//后台登录用户信息
type loginMap struct {
	UserId    int
	User      string
	Ip        string
	LoginTime string
}

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
	id := captcha.New()
	tpl.Execute(w, id)
}

func DoLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		user := r.Form.Get("username")
		pass := r.Form.Get("password")
		captchaCode := r.Form.Get("captcha")
		// 验证码
		captchaId := r.Form.Get("captchaId")

		ok := captcha.VerifyString(captchaId, captchaCode)
		if !ok {
			w.WriteHeader(http.StatusFound)
			w.Write([]byte("验证码错误！"))
			return
		}

		pass += connect.Salt
		id,checkd := dao.AdminLogin(user, utils.Md5(pass))
		if checkd != nil {
			w.WriteHeader(http.StatusFound)
			w.Write([]byte("用户名或密码错误！"))
			return
		}
		// 登陆成功
		ip := utils.ClientIP(r)
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)

		loginInfo := loginMap{id,user, ip, time.Now().Format("2006-01-02 15:04:05")}
		err := sess.Set("userInfo", loginInfo)

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

	http.Redirect(w, r, "/", http.StatusFound)
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

// 文件上传
func UploadFile(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var response []byte
		w.Header().Set("Content-Type", "application/json")
		file, header, err := r.FormFile("file")
		if err != nil {
			response = utils.JsonReturn(connect.ERR_API, "No file uploaded！")
			w.Write(response)
			return
		}
		defer file.Close()

		// 文件上传位置存储|建议OSS
		currDay := time.Now().Format("20060102")
		filePath := "./static/upload/" + currDay + "/"
		exists, err := utils.PathExists(filePath)
		if !exists {
			// 文件夹不存在则新建
			if err = os.Mkdir(filePath, os.ModePerm); err != nil {
				response = utils.JsonReturn(connect.ERR_API, "mkdir failed:"+err.Error())
				w.Write(response)
				return
			}
		}

		newFile, err := os.Create(filePath + header.Filename)
		if err != nil {
			response = utils.JsonReturn(connect.ERR_API, err.Error())
			w.Write(response)
			return
		}
		defer newFile.Close()

		// 移动目标文件路径
		_, err = io.Copy(newFile, file)
		if err != nil {
			response = utils.JsonReturn(connect.ERR_API, "file upload failed")
			w.Write(response)
			return
		}
		newFile.Seek(0, 0)

		response = utils.JsonReturn(connect.OK_API, "file upload success")
		w.Write(response)
	}

}
