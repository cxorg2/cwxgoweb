package blog

<<<<<<< HEAD
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WorkDB struct {
	Dsn string
}

func (d *WorkDB) GetSession() *gorm.DB {
	db, err := gorm.Open(mysql.Open(d.Dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func (d *WorkDB) GetDocPv(url string) int {
	session := d.GetSession()
=======
import dbmysql "git.services.wait/chenwx/cwxgoweb/src/common/dbmysql"

type docDB struct {
	db *dbmysql.MysqlConf
}

func (d *docDB) GetDocPv(url string) int {
	session := d.db.GetSession()
>>>>>>> main
	var doc T_doc_access

	return doc.getPv(session, url)
}
