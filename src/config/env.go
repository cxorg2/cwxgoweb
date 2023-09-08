package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

func GetEnvModuleConfig(g_conf *GlobalConf) {
	getStressConfig(&g_conf.ConfStress)
	getWebServerConf(&g_conf.ConfWebServer)
	getGinWebConf(&g_conf.ConfGinWeb)
	getLogTaskConf(g_conf.ConfLogTask)
	getMsqlConf(&g_conf.ConfMysql, &g_conf.ConfStress)
	getRedisConfig(&g_conf.ConfRedis, &g_conf.ConfStress)
	getMetricConf(&g_conf.ConfMetrics)

}

func getWebServerConf(conf *ConfWebServer) {
	if unit.IsTrue(os.Getenv("CWX_WEBSERVER_ON")) {
		conf.On = true
	} else {
		return
	}
	var err error
	port_str := os.Getenv("CWX_WEBSERVER_PORT")
	conf.Port, err = strconv.Atoi(port_str)
	if err != nil {
		conf.Port = 19002
		log.Println("no webserver port env")
	}

}

func getGinWebConf(conf *ConfGinWeb) {
	var err error
	if unit.IsTrue(os.Getenv("CWX_GINWEB_ON")) {
		conf.On = true
	} else {
		return
	}

	port_str := os.Getenv("CWX_GINWEB_PORT")
	conf.Port, err = strconv.Atoi(port_str)
	if err != nil {
		conf.Port = 19004
		log.Println("no ginweb port env")
	}

}

func getLogTaskConf(conf ConfLogTask) {
	if unit.IsTrue(os.Getenv("CWX_LOGTASK_ON")) {
		conf.On = true
	}
}

func getMsqlConf(conf *ConfMysql, stress *ConfStress) {

	if unit.IsTrue(os.Getenv("CWX_GENERATEDATA_MYSQL_ON")) {
		conf.On = true
		conf.ConfSourcce = "env"
	} else {
		return
	}

	db_addr := os.Getenv("CWX_DB_ADDR")
	db_port := os.Getenv("CWX_DB_PORT")
	db_user := os.Getenv("CWX_DB_USER")
	db_pw := os.Getenv("CWX_DB_PASSWORD")
	db_name := os.Getenv("CWX_DB_NAME")
	db_charset := os.Getenv("CWX_DB_CHARSET")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		db_user, db_pw, db_addr, db_port, db_name, db_charset)

	conf.Dsn = dsn
	conf.SleepMs = stress.MysqlSleepMs

}

func getMetricConf(conf *ConfMetrics) {

	if unit.IsTrue(os.Getenv("CWX_METRICS_ON")) {
		conf.On = true
		conf.ConfSourcce = "env"
	} else {
		return
	}

	port := os.Getenv("CWX_METRICS_PORT")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("get Metrics port error")
		os.Exit(1)
	}

	conf.MetricsPort = portInt

}

func getRedisConfig(conf *ConfRedis, stress *ConfStress) {
	if unit.IsTrue(os.Getenv("CWX_GENERATEDATA_REDIS_ON")) {
		conf.ConfSourcce = "env"
		conf.On = true
	} else {
		return
	}

	conf.RedisAddress = os.Getenv("CWX_REDIS_ADDRESS")
	conf.RedisPort = os.Getenv("CWX_REDIS_PORT")
	conf.RedisAuthPassword = os.Getenv("CWX_REDIS_PASSWORD")
	conf.SleepMs = stress.CacheSleepMs

}

func getStressConfig(conf *ConfStress) {
	cache_sleep := os.Getenv("CWX_REDIS_SLEEP_MS")
	if len(cache_sleep) > 0 {
		conf.CacheSleepMs, _ = strconv.Atoi(cache_sleep)
	} else {
		conf.CacheSleepMs = 200
	}

	mysql_sleep := os.Getenv("CWX_DB_SLEEP_MS")

	if len(mysql_sleep) > 0 {
		conf.CacheSleepMs, _ = strconv.Atoi(mysql_sleep)
	} else {
		conf.MysqlSleepMs = 200
	}

	conf.ConfSourcce = "env"

}
