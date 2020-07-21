package dao

import (
	"database/sql"
	"errors"
	"github.com/go_admin/connect"
	"github.com/go_admin/utils"
	"log"
	"net/http"
	"strings"
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

func AdminPassChange(r *http.Request) error {
	old_pass := strings.TrimSpace(r.Form.Get("old_pass"))
	user := r.Form.Get("username")
	old_pass += connect.Salt
	ok := AdminLogin(user, utils.Md5(old_pass))
	if !ok {
		return errors.New("原密码错误!")
	}

	pass := strings.TrimSpace(r.Form.Get("pass"))
	repass := r.Form.Get("repass")
	if pass != repass {
		return errors.New("两次密码输入不一致!")
	}

	pass_new := utils.Md5(repass + connect.Salt)
	stmt, _ := db.Prepare("UPDATE `go_admin` SET pass = ? WHERE name = ?")
	_, err := stmt.Exec(pass_new, user)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
