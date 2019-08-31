package handler

import (
	"fmt"
	"go_admin/connect"
	"go_admin/dao"
	"go_admin/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

func CateDelHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	cateId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("ID type conversion error")
		return
	}
	err = dao.DelCateByID(cateId)
	var res []byte
	if err != nil {
		res = utils.JsonReturn(connect.ERR_API, "删除失败!")
	} else {
		res = utils.JsonReturn(connect.OK_API, "删除成功!")
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(res)))
	w.Write(res)
}
