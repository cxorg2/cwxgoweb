package generatedata

import (
	"log"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/config"
	"git.services.wait/chenwx/cwxgoweb/src/metrics"
	"git.services.wait/chenwx/cwxgoweb/src/unit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysql insert number
var insertCount int64

type HostInfo struct {
	// gorm.Model
	Id         int       `gorm:"primary_key"`
	Itime      time.Time `gorm:"autoCreateTime;"`
	Ctime      time.Time `gorm:"autoCreateTime;not null;"`
	Ip         string    `gorm:"not null;"`
	Cpu        float64   `gorm:"not null;"`
	Random_num int64     `gorm:"not null;"`
	Status     int       `gorm:"not null;"`
}

func (HostInfo) TableName() string {
	return "t_host_info"
}

func GetSession(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func getOneData() HostInfo {
	currTime := time.Now()
	oldTime := currTime.Add(-time.Hour * 1) // 一个小时前

	var data = HostInfo{
		Ctime:      oldTime,
		Ip:         unit.GetlocalIP(),
		Cpu:        unit.RandCpu(),
		Random_num: unit.RandNumInt64Length(50000),
		Status:     0,
	}
	return data
}

func MysqlTask(conf config.ConfMysql) {

	db := GetSession(conf.Dsn)
	log.Println("task: --- generate data mysql start")

	// 定时计数器
	go unit.CountNumTicker("mysql", &insertCount)

	for {
		data := getOneData()
		db.Create(&data)

		time.Sleep(time.Duration(conf.SleepMs))
		insertCount += 1
		metrics.InsertDB.Inc()
		// insertDB.Inc()
	}

}
