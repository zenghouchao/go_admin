package dao

import (
	"database/sql"
	"github.com/go_admin/connect"
)

var db *sql.DB

func init() {
	db = connect.ConnectDB()
}
