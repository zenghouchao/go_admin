package dao

import (
	"database/sql"
	"go_admin/connect"
)

var db *sql.DB

func init() {
	db = connect.ConnectDB()
}
