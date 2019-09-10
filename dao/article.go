package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go_admin/connect"
	"log"
	"strconv"
	"strings"
)

func AddCategory(cateName string) error {
	stmt, _ := db.Prepare("SELECT id FROM `go_cate` WHERE name = ? ")
	var id string
	err := stmt.QueryRow(cateName).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf(err.Error())
		return err
	}

	cateId, _ := strconv.Atoi(id)

	if cateId > 0 {
		log.Printf("cate name exists")
		return errors.New("cate name exists")
	}

	stmt, err = db.Prepare("INSERT INTO `go_cate` (name) VALUES(?) ")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(cateName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer stmt.Close()
	return nil
}

func DelArticleByID(articleId int) error {
	stmt, _ := db.Prepare("DELETE FROM `go_article` WHERE id = ?")
	_, err := stmt.Exec(articleId)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func GetCateList(cate string, page int) (int, []*connect.Cate, error) {
	cate_sql := ""
	if cate != "" {
		cate_sql += " WHERE name LIKE '" + strings.TrimSpace(cate) + "%'"
	}
	var (
		count int
		res   []*connect.Cate
	)
	err := db.QueryRow(`select count(*) FROM go_cate` + cate_sql).Scan(&count)
	if err != nil {
		return count, res, err
	}

	stmt, err := db.Prepare("SELECT * FROM `go_cate`" + cate_sql + " ORDER BY id DESC LIMIT ?,?")

	rows, err := stmt.Query((page-1)*connect.PageSize, connect.PageSize)

	if err != nil {
		return count, res, err
	}

	for rows.Next() {
		var id, name, status string
		if err := rows.Scan(&id, &name, &status); err != nil {
			return count, res, err
		}
		c := &connect.Cate{
			Id:     id,
			Name:   name,
			Status: status,
		}
		res = append(res, c)
	}
	defer stmt.Close()
	return count, res, nil
}

func DelCateByID(id int) error {
	// 栏目下存在文章则不删除
	stmt, _ := db.Prepare("SELECT id FROM `go_article` WHERE cateId = ?")
	row := stmt.QueryRow(id)
	var articleId string
	if err := row.Scan(&articleId); err != nil && err != sql.ErrNoRows {
		return err
	}

	if articleId != "" {
		return errors.New("该分类下存在文章不能删除!")
	}

	stmt, err := db.Prepare("DELETE FROM `go_cate` WHERE id = ?")
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func SaveCateStatus(id int, status int) error {
	stmt, err := db.Prepare("UPDATE `go_cate` SET `status` = ? WHERE id = ? ")
	_, err = stmt.Exec(status, id)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func AddArticle(dataMap *connect.Article) error {

	stmt, err := db.Prepare("INSERT INTO `go_article` (cateId,title,content,time,status,author) VALUES(?,?,?,?,?,?) ")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(dataMap.Cate_id, dataMap.Title, dataMap.Content, dataMap.Time, dataMap.Status, dataMap.Author)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func UpdateArtice(dataMap *connect.Article) error {
	update_sql := `UPDATE go_article SET cateId =?, title =?, content =?,time=?,
	status=?,author=? WHERE id=?`
	stmt, err := db.Prepare(update_sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(dataMap.Cate_id, dataMap.Title, dataMap.Content, dataMap.Time, dataMap.Status, dataMap.Author, dataMap.Id)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func GetArticleList(page int) (int, []*connect.Article, error) {

	var (
		count  int
		result []*connect.Article
	)
	err := db.QueryRow(`select count(*) FROM go_article AS a 
		LEFT JOIN go_cate AS c ON c.id = a.cateId WHERE c.status = 1`).Scan(&count)
	if err != nil {
		return count, result, err
	}

	stmt, _ := db.Prepare(`SELECT a.id,c.name,a.title,a.content,a.time,a.status,a.author FROM go_article AS a 
		LEFT JOIN go_cate AS c ON c.id = a.cateId WHERE c.status = 1 ORDER BY a.id DESC LIMIT ?,? `)

	rows, err := stmt.Query((page-1)*connect.PageSize, connect.PageSize)

	if err != nil {
		return count, result, err
	}

	for rows.Next() {
		var (
			id, cate, title, content, status, author string
			pubTime                                  int64
		)

		err = rows.Scan(&id, &cate, &title, &content, &pubTime, &status, &author)

		if err != nil {
			return count, result, err
		}

		r := &connect.Article{
			Id:      id,
			Cate_id: cate,
			Title:   title,
			Content: content,
			Time:    pubTime,
			Status:  status,
			Author:  author,
		}
		result = append(result, r)
	}
	defer rows.Close()
	return count, result, nil
}

func GetFirstArticleByID(id int) (*connect.Article, error) {
	stmt, _ := db.Prepare("SELECT * FROM `go_article` WHERE id = ?")
	var (
		aid, cate, title, content, status, author string
		pubTime                                   int64
		r                                         *connect.Article
	)

	err := stmt.QueryRow(id).Scan(&aid, &cate, &title, &content, &pubTime, &status, &author)
	if err != nil && err != sql.ErrNoRows {
		return r, err
	}
	defer stmt.Close()

	r = &connect.Article{
		Id:      aid,
		Cate_id: cate,
		Title:   title,
		Content: content,
		Time:    pubTime,
		Status:  status,
		Author:  author,
	}
	return r, nil

}

func GetOnCate() ([]*connect.Cate, error) {
	cate_sql := "SELECT * FROM `go_cate` WHERE status = 1 "

	stmt, err := db.Prepare(cate_sql + " ORDER BY id DESC LIMIT ?")
	var res []*connect.Cate
	rows, err := stmt.Query(connect.PageSize)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, status string
		if err := rows.Scan(&id, &name, &status); err != nil {
			return res, err
		}
		c := &connect.Cate{
			Id:     id,
			Name:   name,
			Status: status,
		}
		res = append(res, c)
	}
	defer stmt.Close()
	return res, nil
}
