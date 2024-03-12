package blog

import (
	"fmt"
	"os"

	"git.services.wait/chenwx/cwxgoweb/src/common/dbmysql"
	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

type BlogMysql struct {
	Enable bool
	*dbmysql.MysqlConf
}

// BLOG
type BlogConf struct {
	Enable bool
	Port   string
	Mysql  BlogMysql
}

func (C *BlogConf) GetEnvConf() {
	if unit.IsTrue(os.Getenv("CWX_BLOG_API_ENABLE")) {
		C.Enable = true
		port := os.Getenv("CWX_BLOG_API_PORT")
		if port == "" {
			port = "19001"
		}
		C.Port = port
	}

	if unit.IsTrue(os.Getenv("CWX_BLOG_MYSQL_ENABLE")) {
		C.Mysql.Enable = true
		var m dbmysql.MysqlConf
		C.Mysql.MysqlConf = &m

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
