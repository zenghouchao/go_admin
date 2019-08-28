package handler

import (
	"fmt"
	"go_admin/dao"
	"go_admin/utils"
	"html/template"
	"io"
	"net/http"
	"os"
)

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/article/list.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, nil)
}

func ArticleAddHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/article/add.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, nil)
}

func CateListHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/article/cate.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, nil)
}

func CateAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		cateName := r.Form.Get("cate")
		var result []byte

		err = dao.AddCategory(cateName)
		if err != nil {
			result = utils.JsonReturn(1, "新增栏目失败")
		} else {
			result = utils.JsonReturn(0, "新增栏目成功")
		}
		io.WriteString(w, string(result))
		os.Exit(0)
	}
	tpl, err := template.ParseFiles("./template/article/addCate.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, nil)
}
