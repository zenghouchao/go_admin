package connect

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func ConnectDB() *sql.DB {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go_web?charset=utf8")

	if err != nil {
		panic(err.Error())
	}
	return db
}
