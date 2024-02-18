package dbmysql

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConf struct {
	Address  string
	User     string
	Password string
	Dbname   string
	Port     string
	Dsn      string
	CharSet  string
}

func (C *MysqlConf) GetSession() *gorm.DB {
	db, err := gorm.Open(mysql.Open(C.Dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

// 非 orm, 原始SQL
func (C *MysqlConf) GetSessionSqlDB(dsn string) *sql.DB {
	// "wait:passw0rd@tcp(39.106.0.117:39008)/mytest?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
