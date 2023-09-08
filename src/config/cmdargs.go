package config

import (
	"flag"
	"log"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// 解析本地参数
func appArgsConf(g_conf *GlobalConf) {
	var stressSize int
	var localAddr string

	flag.IntVar(&stressSize, "p", 0, "pressure 1: low, 2: medium, 3: high; default 0 get consul")
	flag.StringVar(&localAddr, "n", "", "local ipaddress, default: 10.x.x.x")
	flag.Parse()

	// 有设置
	if len(localAddr) != 0 {
		localAddr = unit.GetlocalIP()
		g_conf.ConfigLocal.localAddr = localAddr
		g_conf.ConfigLocal.ConfSourcce = "cli"
	} else {
		g_conf.ConfigLocal.ConfSourcce = "env"
	}

	// 如果没有配置 modelSize, 则使用 env 配置
	if stressSize == 0 {
		log.Println("arg modelSize 0; use env conf")
		return
	}

	setStress(stressSize, g_conf)

}

func setStress(stressSize int, g_conf *GlobalConf) {

	if stressSize == 1 {
		g_conf.ConfStress.ConfSourcce = "cli"

		g_conf.ConfStress.CacheSleepMs = 200
		g_conf.ConfRedis.SleepMs = 200

		g_conf.ConfStress.MysqlSleepMs = 900
		g_conf.ConfMysql.SleepMs = 900
	}

	if stressSize == 2 {
		g_conf.ConfStress.ConfSourcce = "cli"
		g_conf.ConfStress.CacheSleepMs = 50
		g_conf.ConfRedis.SleepMs = 50

		g_conf.ConfStress.MysqlSleepMs = 300
		g_conf.ConfMysql.SleepMs = 300
	}

	if stressSize >= 3 {
		g_conf.ConfStress.ConfSourcce = "cli"
		g_conf.ConfStress.CacheSleepMs = 20
		g_conf.ConfRedis.SleepMs = 20
		g_conf.ConfStress.MysqlSleepMs = 100
		g_conf.ConfMysql.SleepMs = 100
	}

}
