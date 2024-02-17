package generatedata

import (
	"log"
	"strconv"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/metrics"
	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// mysql insert number
var insertCount int64

type hostInfo struct {
	// gorm.Model
	Id         int       `gorm:"primary_key"`
	Itime      time.Time `gorm:"autoCreateTime;"`
	Ctime      time.Time `gorm:"autoCreateTime;not null;"`
	Ip         string    `gorm:"not null;"`
	Cpu        float64   `gorm:"not null;"`
	Random_num int64     `gorm:"not null;"`
	Status     int       `gorm:"not null;"`
}

func (hostInfo) TableName() string {
	return "t_host_info"
}

// func getSession(dsn string) *gorm.DB {
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}

// 	return db
// }

func getOneData() hostInfo {
	currTime := time.Now()
	oldTime := currTime.Add(-time.Hour * 1) // 一个小时前

	var data = hostInfo{
		Ctime:      oldTime,
		Ip:         unit.GetlocalIP(),
		Cpu:        unit.RandCpu(),
		Random_num: unit.RandNumInt64Length(50000),
		Status:     0,
	}
	return data
}

func mysqlTask(cfg GenerateConf) {
	db := cfg.Mysql.MysqlConf.GetSession()

	log.Println("generatedata: to mysql start")

	// 定时计数器
	go unit.CountNumTicker("mysql", &insertCount)

	for {
		data := getOneData()
		db.Create(&data)

		i, _ := strconv.ParseInt(cfg.Mysql.SleepMs, 10, 64)
		time.Sleep(time.Duration(i))
		insertCount += 1
		metrics.InsertDB.Inc()
		// insertDB.Inc()
	}

}
