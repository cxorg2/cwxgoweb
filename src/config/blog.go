package config

import (
	"fmt"
	"os"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// BLOG
type Blog struct {
	Api struct {
		Enable bool
		Port   string
	}

	Mysql struct {
		Enable   bool
		Address  string
		User     string
		Password string
		Dbname   string
		Port     string
		Dsn      string
		CharSet  string
	}
}

func (C *Blog) getConf() {
	if unit.IsTrue(os.Getenv("CWX_BLOG_API_ENABLE")) {
		C.Api.Enable = true
		port := os.Getenv("CWX_BLOG_API_PORT")
		if port == "" {
			port = "19001"
		}
		C.Api.Port = port
	}

	if unit.IsTrue(os.Getenv("CWX_BLOG_MYSQL_ENABLE")) {
		C.Mysql.Enable = true

		m := &C.Mysql

		m.Enable = true
		m.Address = os.Getenv("CWX_BLOG_MYSQL_ADDR")
		m.Port = os.Getenv("CWX_BLOG_MYSQL_PORT")
		m.User = os.Getenv("CWX_BLOG_MYSQL_USER")
		m.Password = os.Getenv("CWX_BLOG_MYSQL_PW")
		m.Dbname = os.Getenv("CWX_BLOG_MYSQL_DBNAME")
		m.CharSet = os.Getenv("CWX_BLOG_MYSQL_CHARSET")
		m.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
			m.User, m.Password, m.Address, m.Port, m.Dbname, m.CharSet)

	}

}
