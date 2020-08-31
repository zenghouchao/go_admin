package handler

import (
	"fmt"
	"github.com/go_admin/connect"
	"github.com/go_admin/dao"
	pager "github.com/go_admin/page"
	"github.com/go_admin/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	var p int
	query := r.URL.Query()
	page := query.Get("page")

	if page == "" {
		p = 1
		page = "1"
	} else {
		p, _ = strconv.Atoi(page)
	}

	// get Article info
	count, data, err := dao.GetArticleList(p)
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
	// page
	ps := pager.NewUrlPager(pager.MathPages(count, connect.PageSize), p, "/article?page=%d")
	pagerHtml := ps.PagerString()

	params := map[string]interface{}{
		"page": template.HTML(pagerHtml),
		"list": data,
	}

	w.Header().Set("Content-Type", "text/html")
	if err = tpl.Execute(w, params); err != nil {
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
			Title:   utils.InputSafe(r.PostForm.Get("title")),
			Content: utils.InputSafe(r.PostForm.Get("desc")),
			Time:    theTime.Unix(),
			Status:  r.PostForm.Get("status"),
			Author:  utils.InputSafe(r.PostForm.Get("author")),
		}
		err := dao.AddArticle(data)
		var result []byte
		if err != nil {
			result = utils.JsonReturn(connect.ERR_API, "Failed to publish article")
		} else {
			result = utils.JsonReturn(connect.OK_API, "Success to publish article")
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(result)))
		w.Write(result)

	} else {

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
			result = utils.JsonReturn(connect.ERR_API, "Article ID cannot be empty")
			w.Write(result)
			return
		}
		id, _ := strconv.Atoi(articleId)
		err := dao.DelArticleByID(id)
		if err != nil {
			result = utils.JsonReturn(connect.ERR_API, "Failed to delete article")
		} else {
			result = utils.JsonReturn(connect.OK_API, "Success to delete article")
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
			Title:   utils.InputSafe(postFrom.Get("title")),
			Content: utils.InputSafe(postFrom.Get("desc")),
			Time:    theTime.Unix(),
			Status:  postFrom.Get("status"),
			Author:  utils.InputSafe(postFrom.Get("author")),
		}

		err := dao.UpdateArtice(data)
		if err != nil {
			response = utils.JsonReturn(connect.ERR_API, "Failed！")
		} else {
			response = utils.JsonReturn(connect.ERR_API, "Success！")
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(response)))
		w.Write(response)
	}
}

// article edit page
func ArticleEditPageHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	aid, _ := strconv.Atoi(id)
	// Get article info
	data, err := dao.GetFirstArticleByID(aid)
	if err != nil {
		fmt.Println("get article data failure:" + err.Error())
	}
	formatDate := time.Unix(data.Time, 0).Format("2006-01-02")
	data.Pubdate = formatDate

	// Get column data
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
	var p int
	tpl, err := template.ParseFiles("./template/article/cate.html")
	if err != nil {
		fmt.Println("Loading template error:" + err.Error())
		return
	}
	// search request
	query := r.URL.Query()
	catName := query.Get("cate_name")
	page := query.Get("page")

	if page == "" {
		p = 1
	} else {
		p, _ = strconv.Atoi(page)
	}

	// Get column data
	count, data, catesErr := dao.GetCateList(catName, p)
	if catesErr != nil {
		fmt.Println("Get column data Fail")
		return
	}

	// page
	ps := pager.NewUrlPager(pager.MathPages(count, connect.PageSize), p, "/cate?page=%d")
	pagerHtml := ps.PagerString()

	params := map[string]interface{}{
		"page": template.HTML(pagerHtml),
		"list": data,
	}

	w.Header().Set("Content-Type", "text/html")
	if err = tpl.Execute(w, params); err != nil {
		fmt.Println("load cate template error:", err.Error())
	}
}

func CateAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		cateName := r.Form.Get("cate")
		var result []byte

		if err := dao.AddCategory(cateName); err != nil {
			result = utils.JsonReturn(connect.ERR_API, "Failed")
		} else {
			result = utils.JsonReturn(connect.OK_API, "Success")
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
			res = utils.JsonReturn(connect.OK_API, "Success!")
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
			res = utils.JsonReturn(connect.ERR_API, "Missing parameter error!")
			w.Write(res)
			return
		}
		cateStatus, _ := strconv.Atoi(status)
		cateId, _ := strconv.Atoi(id)

		err := dao.SaveCateStatus(cateId, cateStatus)
		if err != nil {
			res = utils.JsonReturn(connect.ERR_API, "Status update failed!")
		} else {
			res = utils.JsonReturn(connect.OK_API, "Status update successful!")
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(res)))
		w.Write(res)
	}

}
