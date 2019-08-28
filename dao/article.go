package dao

import (
	"database/sql"
	"fmt"
	"go_admin/connect"
)

func AddCategory(cateName string) error {
	db := connect.Init()

	stmt, err := db.Prepare("INSERT INTO `go_cate` (name) VALUES(?) ")

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
