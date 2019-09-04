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
	"strings"
	"time"
)

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	// 获取文章信息
	data, err := dao.GetArticleList()
	if err != nil {
		fmt.Println("get article data failure", err.Error())
		return
	}
	for _, item := range data {
		pubdate := time.Unix(item.Time, 0).Format("2006-01-02")
		item.Pubdate = pubdate
	}
	tpl, err := template.ParseFiles("./template/article/list.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	if err = tpl.Execute(w, data); err != nil {
		panic(err.Error())
	}
}

func ArticleAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		pubdate := r.PostForm.Get("pubdate")
		loc, _ := time.LoadLocation("Local")
		theTime, _ := time.ParseInLocation("2006-01-02", pubdate, loc)

		data := &connect.Article{
			Cate_id: r.PostForm.Get("cateId"),
			Title:   strings.TrimSpace(r.PostForm.Get("title")),
			Content: strings.TrimSpace(r.PostForm.Get("desc")),
			Time:    theTime.Unix(),
			Status:  r.PostForm.Get("status"),
			Author:  strings.TrimSpace(r.PostForm.Get("author")),
		}
		err := dao.AddArticle(data)
		var result []byte
		if err != nil {
			result = utils.JsonReturn(connect.ERR_API, "发布文章失败")
		} else {
			result = utils.JsonReturn(connect.OK_API, "发布文章成功")
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(result)))
		w.Write(result)

	} else {
		// 获取所以栏目
		cates, err := dao.GetOnCate()
		if err != nil {
			log.Println("no found cate data error:", err.Error())
		}

		tpl, err := template.ParseFiles("./template/article/add.html")
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
}

func ArticleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var result []byte
	if r.Method == "POST" {
		r.ParseForm()
		articleId := r.Form.Get("id")
		if articleId == "" {
			result = utils.JsonReturn(connect.ERR_API, "文章ID不能为空")
			w.Write(result)
			return
		}
		id, _ := strconv.Atoi(articleId)
		err := dao.DelArticleByID(id)
		if err != nil {
			result = utils.JsonReturn(connect.ERR_API, "删除文章失败")
		} else {
			result = utils.JsonReturn(connect.OK_API, "删除文章成功")
		}

		w.Header().Set("Content-Length", strconv.Itoa(len(result)))
		w.Write(result)
	}
}

func ArticleUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var response []byte
	if r.Method == "POST" {
		r.ParseForm()
		postFrom := r.PostForm

		pubdate := postFrom.Get("pubdate")
		loc, _ := time.LoadLocation("Local")
		theTime, _ := time.ParseInLocation("2006-01-02", pubdate, loc)

		data := &connect.Article{
			Id:      postFrom.Get("id"),
			Cate_id: postFrom.Get("cateId"),
			Title:   strings.TrimSpace(postFrom.Get("title")),
			Content: strings.TrimSpace(postFrom.Get("desc")),
			Time:    theTime.Unix(),
			Status:  postFrom.Get("status"),
			Author:  strings.TrimSpace(postFrom.Get("author")),
		}

		err := dao.UpdateArtice(data)
		if err != nil {
			response = utils.JsonReturn(connect.ERR_API, "更新文章失败！")
		} else {
			response = utils.JsonReturn(connect.ERR_API, "更新文章成功！")
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(response)))
		w.Write(response)
	}
}

// 编辑文章页面
func ArticleEditPageHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	aid, _ := strconv.Atoi(id)
	// 获取文章信息
	data, err := dao.GetFirstArticleByID(aid)
	if err != nil {
		fmt.Println("get article data failure:" + err.Error())
	}
	formatDate := time.Unix(data.Time, 0).Format("2006-01-02")
	data.Pubdate = formatDate

	// 获取栏目
	cates, err := dao.GetOnCate()
	if err != nil {
		log.Println("no found cate data error:", err.Error())
	}

	params := map[string]interface{}{
		"cates":   cates,
		"article": data,
	}

	tpl, err := template.ParseFiles("./template/article/edit.html")
	if err != nil {
		fmt.Println("Loading article edit page template error:" + err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	if err = tpl.Execute(w, params); err != nil {
		fmt.Println("load article edit page template failure:", err.Error())
	}
}

func CateListHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/article/cate.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	// 搜索请求
	query := r.URL.Query()
	catName := query.Get("cate_name")

	// 获取栏目数据
	res, catesErr := dao.GetCateList(catName)
	if catesErr != nil {
		fmt.Println("获取栏目数据失败")
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err = tpl.Execute(w, res); err != nil {
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
			res = utils.JsonReturn(connect.ERR_API, err.Error())
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
