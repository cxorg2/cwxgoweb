package dbmysql

import (
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
