package generatedata

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/common/dbmysql"
	"git.services.wait/chenwx/cwxgoweb/src/common/redis"
	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

type stressRedis struct {
	Enable  bool
	SleepMs string
	*redis.RedisConf
}

type stressMysql struct {
	Enable  bool
	SleepMs string
	*dbmysql.MysqlConf
}

// 压力模块配置
type GenerateConf struct {
	Enable bool
	Speed  int // 预设压力大小;0 默认值, 1: low, 2: medium, 3: high
	Redis  stressRedis
	Mysql  stressMysql
}

// 获取压力配置
func (C *GenerateConf) GetEnvConf() {

	stressEnable := os.Getenv("CWX_GENERATEDATA_ENABLE")

	if unit.IsTrue(stressEnable) {
		C.Enable = true
	} else {
		return
	}

	if unit.IsTrue(os.Getenv("CWX_GENERATEDATA_REDIS_ENABLE")) {

		C.Redis.Enable = true
		C.Redis.SleepMs = os.Getenv("CWX_GENERATEDATA_REDIS_SLEEP_MS")

		var r redis.RedisConf
		r.Address = os.Getenv("CWX_GENERATEDATA_REDIS_ADDR")
		r.Port = os.Getenv("CWX_GENERATEDATA_REDIS_PORT")
		r.AuthPassword = os.Getenv("CWX_GENERATEDATA_REDIS_PW")

		C.Redis.RedisConf = &r

	}

	if unit.IsTrue(os.Getenv("CWX_GENERATEDATA_MYSQL_ENABLE")) {

		C.Mysql.Enable = true
		C.Mysql.SleepMs = os.Getenv("CWX_GENERATEDATA_MYSQL_SLEEP_MS")

		var m dbmysql.MysqlConf

		m.Address = os.Getenv("CWX_GENERATEDATA_MYSQL_ADDR")
		m.Port = os.Getenv("CWX_GENERATEDATA_MYSQL_PORT")
		m.User = os.Getenv("CWX_GENERATEDATA_MYSQL_USER")
		m.Password = os.Getenv("CWX_GENERATEDATA_MYSQL_PW")
		m.Dbname = os.Getenv("CWX_GENERATEDATA_MYSQL_DBNAME")
		m.CharSet = os.Getenv("CWX_GENERATEDATA_MYSQL_CHARSET")
		m.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
			m.User, m.Password, m.Address, m.Port, m.Dbname, m.CharSet)

		C.Mysql.MysqlConf = &m

	}
	// Speed speed
	speedSize := os.Getenv("CWX_GENERATEDATA_SPEED")

	if unit.IsTrue(speedSize) {
		int_level, _ := strconv.Atoi(speedSize)
		if int_level != 0 {
			setSpeed(int_level, C)
		}

	}

}

func (C *GenerateConf) GetCmdArgs() {
	var speedSize int
	flag.IntVar(&speedSize, "p", 0, "pressure 1: low, 2: medium, 3: high; default 0 used env")

	// 如果有配置 modelSize, 则使用 env 配置
	if speedSize != 0 {
		setSpeed(speedSize, C)
		log.Printf("generatedata: arg modelSize %d; use env conf\n", speedSize)
		return
	}

}

func setSpeed(speedSize int, C *GenerateConf) {

	if speedSize == 1 {
		C.Redis.SleepMs = "200"
		C.Mysql.SleepMs = "900"
	}

	if speedSize == 2 {
		C.Redis.SleepMs = "100"
		C.Mysql.SleepMs = "500"
	}

	if speedSize >= 3 {
		C.Redis.SleepMs = "50"
		C.Mysql.SleepMs = "100"
	}

}

func RunGenerateData(cfg GenerateConf) {
	if !cfg.Enable {
		log.Println("model: no enable GenerateData")
		return
	}

	if cfg.Redis.Enable {
		go redisTask(cfg)
	}

	if cfg.Mysql.Enable {
		go mysqlTask(cfg)
	}

}
