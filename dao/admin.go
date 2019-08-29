package dao

import (
	"database/sql"
	"log"
)

func AdminLogin(user string, pass string) bool {
	stmt, err := db.Prepare("SELECT pass FROM `go_admin` WHERE name = ?")
	if err != nil {
		log.Printf("%s\n", err)
		return false
	}
	var pwd string
	err = stmt.QueryRow(user).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return false
	}

	defer stmt.Close()
	if pass != pwd {
		return false
	}
	return true
}
