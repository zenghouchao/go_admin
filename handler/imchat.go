package handler

import (
	"fmt"
	"github.com/go_admin/connect"
	"github.com/go_admin/dao"
	"github.com/go_admin/utils"
	"net/http"
	"html/template"
	"strconv"
)

func ImChatHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pageStr := query.Get("page")
	var p int
	if err != nil {
		panic(err)
	}
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

	tpl, err := template.ParseFiles("./template/chat/index.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")

	var params map[string]interface{}
	params = map[string]interface{}{
		"cates": cates,
	}
	if err = tpl.Execute(w, params); err != nil {
		fmt.Printf("add article template load error: ", err.Error())
	}
}
