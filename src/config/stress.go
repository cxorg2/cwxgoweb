package config

import (
	"fmt"
	"os"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

type StressRedis struct {
	Enable  bool
	SleepMs string
	*Redis
}

type StressMysql struct {
	Enable  bool
	SleepMs string
	*Mysql
}

// 压力模块配置
type Stress struct {
	Enable     bool
	StressSize int // 预设压力大小;0 默认值, 1: low, 2: medium, 3: high
	Redis      StressRedis
	Mysql      StressMysql
}

// 获取压力配置
func (C *Stress) getConf() {

	stressEnable := os.Getenv("CWX_STRESS_ENABLE")

	if unit.IsTrue(stressEnable) {
		C.Enable = true
	} else {
		return
	}

	if unit.IsTrue(os.Getenv("CWX_STRESS_REDIS_ENABLE")) {
		C.Redis.Enable = true
		C.Redis.SleepMs = os.Getenv("CWX_STRESS_REDIS_SLEEP_MS")

		var r Redis
		r.Address = os.Getenv("CWX_STRESS_REDIS_ADDR")
		r.Port = os.Getenv("CWX_STRESS_REDIS_PORT")
		r.AuthPassword = os.Getenv("CWX_STRESS_REDIS_PW")

		C.Redis.Redis = &r

	}

	if unit.IsTrue(os.Getenv("CWX_STRESS_MYSQL_ENABLE")) {
		C.Mysql.Enable = true
		C.Mysql.SleepMs = os.Getenv("CWX_STRESS_MYSQL_SLEEP_MS")

		var m Mysql

		m.Address = os.Getenv("CWX_STRESS_MYSQL_ADDR")
		m.Port = os.Getenv("CWX_STRESS_MYSQL_PORT")
		m.User = os.Getenv("CWX_STRESS_MYSQL_USER")
		m.Password = os.Getenv("CWX_STRESS_MYSQL_PW")
		m.Dbname = os.Getenv("CWX_STRESS_MYSQL_DBNAME")
		m.CharSet = os.Getenv("CWX_STRESS_MYSQL_CHARSET")
		m.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
			m.User, m.Password, m.Address, m.Port, m.Dbname, m.CharSet)

		C.Mysql.Mysql = &m
	}

	stressSize := os.Getenv("CWX_STRESS_SIZE")

	if unit.IsTrue(stressSize) {
		int_level, _ := strconv.Atoi(stressSize)
		if int_level != 0 {
			setStress(int_level, C)
		}

	}

}

func setStress(stressSize int, C *Stress) {

	if stressSize == 1 {
		C.Redis.SleepMs = "200"
		C.Mysql.SleepMs = "900"
	}

	if stressSize == 2 {
		C.Redis.SleepMs = "100"
		C.Mysql.SleepMs = "500"
	}

	if stressSize >= 3 {
		C.Redis.SleepMs = "50"
		C.Mysql.SleepMs = "100"
	}

}
