package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"go_admin/connect"
	"log"
	"strconv"
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

func GetCateList() ([]*connect.Cate, error) {
	stmt, err := db.Prepare("SELECT * FROM `go_cate` ORDER BY id DESC LIMIT ?")
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

func DelCateByID(id int) error {
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
