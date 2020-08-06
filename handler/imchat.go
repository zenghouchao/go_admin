package handler

import (
	"fmt"
	"github.com/go_admin/dao"
	"net/http"
	"html/template"
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
	users, err := dao.GetUsers(p)

	for k, v := range users {
		fmt.Println(k, v)
	}
	if err != nil {
		panic(err)
	}

	tpl, tErr := template.ParseFiles("./template/chat/index.html")
	if tErr != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")

	var params map[string]interface{}
	params = map[string]interface{}{
		"users": users,
	}
	if err = tpl.Execute(w, params); err != nil {
		fmt.Printf("add article template load error: ", err.Error())
	}
}
