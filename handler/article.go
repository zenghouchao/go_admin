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
		var result []byte

		if err := dao.AddCategory(cateName); err != nil {
			result = utils.JsonReturn(connect.ERR_API, "新增栏目失败")
		} else {
			result = utils.JsonReturn(connect.OK_API, "新增栏目成功")
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(result)))
		w.Write(result)
	} else {
		tpl, err := template.ParseFiles("./template/article/addCate.html")
		if err != nil {
			fmt.Println("Loading template error:" + err.Error())
			return
		}
		w.Header().Set("Content-Type", "text/html")
		tpl.Execute(w, nil)
	}
}

func CateDelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
}

func CateStatusSaveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		status := r.Form.Get("status")
		id := r.Form.Get("id")

		fmt.Println(status, id)

		var res []byte
		if status == "" || id == "" {
			res = utils.JsonReturn(connect.ERR_API, "缺少参数错误!")
			w.Write(res)
			return
		}
		cateStatus, _ := strconv.Atoi(status)
		cateId, _ := strconv.Atoi(id)

		err := dao.SaveCateStatus(cateId, cateStatus)
		if err != nil {
			res = utils.JsonReturn(connect.ERR_API, "状态更新失败!")
		} else {
			res = utils.JsonReturn(connect.OK_API, "状态更新成功!")
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(res)))
		w.Write(res)
	}

}
