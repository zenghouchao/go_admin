package handler

import (
	"fmt"
	"github.com/go_admin/dao"
	"net/http"
	"html/template"
	"reflect"
	"strconv"
)

func ImChatHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pageStr := query.Get("page")
	var p int
	if pageStr == "" {
		p = 1
	} else {
		p, _ = strconv.Atoi(pageStr)
	}

	// get current user
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	adminInfo := sess.Get("userInfo")
	immutable := reflect.ValueOf(adminInfo)
	userId := immutable.FieldByName("UserId").Int()

	users, err := dao.GetUsers(int(userId), p)

	if err != nil {
		panic(err)
	}

	tpl, tErr := template.ParseFiles("./template/chat/index.html")
	if tErr != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")

	params := map[string]interface{}{
		"users": users,
		"user":  userId,
	}
	if err = tpl.Execute(w, params); err != nil {
		fmt.Printf("add article template load error: ", err.Error())
	}
}
