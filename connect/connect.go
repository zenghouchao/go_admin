package connect

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	ini "gopkg.in/ini.v1"
)

var (
	db  *sql.DB
	err error
)

func ConnectDB() *sql.DB {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err.Error())
	}

	hostname := cfg.Section("mysql").Key("hostname").String()
	database := cfg.Section("mysql").Key("database").String()
	username := cfg.Section("mysql").Key("username").String()
	password := cfg.Section("mysql").Key("password").String()
	charset := cfg.Section("mysql").Key("charset").String()
	port := cfg.Section("mysql").Key("port").String()

	dsn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + database + "?charset=" + charset
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	if db.Ping() != nil {
		panic(err.Error())
	}
	return db
}
