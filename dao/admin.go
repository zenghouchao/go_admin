package dao

import (
	"database/sql"
	"errors"
	"github.com/go_admin/connect"
	"github.com/go_admin/utils"
	"net/http"
	"strings"
)

func AdminLogin(user string, pass string) (int, error) {
	var id int
	stmt, err := db.Prepare("SELECT id,pass FROM `go_admin` WHERE name = ?")
	if err != nil {
		return 0, err
	}
	var pwd string

	err = stmt.QueryRow(user).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return id, err
	}
	defer stmt.Close()
	if pass != pwd {
		return id, err
	}

	return id, nil
}

func AdminPassChange(r *http.Request) error {
	old_pass := strings.TrimSpace(r.Form.Get("old_pass"))
	user := r.Form.Get("username")
	old_pass += connect.Salt
	id, ok := AdminLogin(user, utils.Md5(old_pass))
	if ok != nil {
		return errors.New("Original password error!")
	}

	pass := strings.TrimSpace(r.Form.Get("pass"))
	repass := r.Form.Get("repass")
	if pass != repass {
		return errors.New("The two password inputs are inconsistent\n\n!")
	}

	pass_new := utils.Md5(repass + connect.Salt)
	stmt, _ := db.Prepare("UPDATE `go_admin` SET pass = ? WHERE name = ? AND id = ?")
	_, err := stmt.Exec(pass_new, user, id)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

// Get chat room users
func GetUsers(me int, page int) ([]*connect.User, error) {
	var (
		username string
		id       int
		res      []*connect.User
	)

	stmt, _ := db.Prepare("SELECT id,name FROM `go_admin` WHERE id <> ? ORDER BY id DESC LIMIT ?,?")

	rows, err := stmt.Query(me, (page-1)*connect.PageSize, connect.PageSize)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &username); err != nil {
			return res, err
		}
		s := &connect.User{
			Id:   id,
			Name: username,
		}
		res = append(res, s)
	}
	defer stmt.Close()
	return res, nil
}
