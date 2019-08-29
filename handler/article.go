package handler

import (
	"fmt"
	"go_admin/dao"
	"html/template"
	"net/http"
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
	// 获取栏目数据
	res, err := dao.GetCateList()
	if err != nil {
		fmt.Println("获取栏目数据失败")
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tpl.Execute(w, res); err != nil {
		fmt.Println("load cate template error:", err.Error())
	}
}

func CateAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		cateName := r.Form.Get("cate")
		// var (
		// 	result   []byte
		// 	emptyArr interface{}
		// )

		if err := dao.AddCategory(cateName); err != nil {
			fmt.Println("新增栏目失败")
			return
		}

		http.Redirect(w, r, "/cate", http.StatusFound)
		return
		// if err != nil {
		// 	result = utils.JsonReturn(1, "新增栏目失败", emptyArr)
		// } else {
		// 	result = utils.JsonReturn(0, "新增栏目成功", emptyArr)
		// }
		// w.WriteHeader(201)
		// io.WriteString(w, string(result))

	}
	tpl, err := template.ParseFiles("./template/article/addCate.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, nil)
}
