package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

var GlobalConfData GlobalConf

// 获取全局配置
func GetGlobalConf() GlobalConf {
	GlobalConfData.getStressConfig()

	GlobalConfData.getWebServerConf()

	GlobalConfData.getLogTaskConf()

	GlobalConfData.getGinWebConf()

	GlobalConfData.getMsqlConf()

	GlobalConfData.getRedisConfig()

	GlobalConfData.getBlogApiConf()
	GlobalConfData.getBlogMysqlConf()

	GlobalConfData.getMetricConf()

	// 获取命令行参数, 高优先级覆盖环境变量的配置
	GlobalConfData.appArgsConf()

	return GlobalConfData
}

// 获取压力配置
func (g_conf *GlobalConf) getStressConfig() {

	cache_sleep := os.Getenv("CWX_REDIS_SLEEP_MS")

	if len(cache_sleep) > 0 {
		g_conf.ConfStress.CacheSleepMs, _ = strconv.Atoi(cache_sleep)
	} else {
		g_conf.ConfStress.CacheSleepMs = 200
	}

	g_conf.ConfStress.ConfSourcce = "env"

}

// 获取webserver配置
func (g_conf *GlobalConf) getWebServerConf() {

	if unit.IsTrue(os.Getenv("CWX_WEBSERVER_ENABLE")) {
		g_conf.ConfWebServer.On = true
	} else {
		return
	}

	var err error
	port_str := os.Getenv("CWX_WEBSERVER_PORT")
	g_conf.ConfWebServer.Port, err = strconv.Atoi(port_str)
	if err != nil {
		g_conf.ConfWebServer.Port = 19002
		log.Println("no webserver port env")
	}

}

func (g_conf *GlobalConf) getGinWebConf() {
	var err error
	if unit.IsTrue(os.Getenv("CWX_GINWEB_ENABLE")) {
		g_conf.ConfGinWeb.On = true
	} else {
		return
	}

	port_str := os.Getenv("CWX_GINWEB_PORT")
	g_conf.ConfGinWeb.Port, err = strconv.Atoi(port_str)
	if err != nil {
		g_conf.ConfGinWeb.Port = 19004
		log.Println("no ginweb port env")
	}

}

func (g_conf *GlobalConf) getLogTaskConf() {
	if unit.IsTrue(os.Getenv("CWX_LOGTASK_ON")) {
		g_conf.ConfLogTask.On = true
	}
}

func (g_conf *GlobalConf) getMsqlConf() {

	if unit.IsTrue(os.Getenv("CWX_GENERATEDATA_MYSQL_ENABLE")) {
		g_conf.ConfMysql.On = true
		g_conf.ConfMysql.ConfSourcce = "env"
	} else {
		return
	}

	db_addr := os.Getenv("CWX_DB_ADDR")
	db_port := os.Getenv("CWX_DB_PORT")
	db_user := os.Getenv("CWX_DB_USER")
	db_pw := os.Getenv("CWX_DB_PASSWORD")
	db_name := os.Getenv("CWX_GENERATEDATA_DB_NAME")
	db_charset := os.Getenv("CWX_DB_CHARSET")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		db_user, db_pw, db_addr, db_port, db_name, db_charset)

	g_conf.ConfMysql.Dsn = dsn
	g_conf.ConfMysql.SleepMs = g_conf.ConfStress.MysqlSleepMs

}

func (g_conf *GlobalConf) getBlogApiConf() {

	if unit.IsTrue(os.Getenv("CWX_BLOG_API_ENABLE")) {
		g_conf.ConfBlogApi.On = true
		g_conf.ConfBlogApi.ConfSourcce = "env"
	} else {
		return
	}

	port := os.Getenv("CWX_BLOG_API_PORT")
	if port == "" {
		port = "19001"
	}

	g_conf.ConfBlogApi.Port = port

}

func (g_conf *GlobalConf) getBlogMysqlConf() {

	if unit.IsTrue(os.Getenv("CWX_BLOG_API_ENABLE")) {
		g_conf.ConfBlogMysql.On = true
		g_conf.ConfBlogMysql.ConfSourcce = "env"
	} else {
		return
	}

	db_addr := os.Getenv("CWX_DB_ADDR")
	db_port := os.Getenv("CWX_DB_PORT")
	db_user := os.Getenv("CWX_DB_USER")
	db_pw := os.Getenv("CWX_DB_PASSWORD")
	db_name := os.Getenv("CWX_BLOG_API_DB_NAME")
	db_charset := os.Getenv("CWX_DB_CHARSET")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		db_user, db_pw, db_addr, db_port, db_name, db_charset)

	g_conf.ConfBlogMysql.Dsn = dsn

}

func (g_conf *GlobalConf) getMetricConf() {

	if unit.IsTrue(os.Getenv("CWX_METRICS_ENABLE")) {
		g_conf.ConfMetrics.On = true
		g_conf.ConfMetrics.ConfSourcce = "env"
	} else {
		return
	}

	port := os.Getenv("CWX_METRICS_PORT")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("get Metrics port error")
		os.Exit(1)
	}

	g_conf.ConfMetrics.MetricsPort = portInt

}

func (g_conf *GlobalConf) getRedisConfig() {
	if unit.IsTrue(os.Getenv("CWX_GENERATEDATA_REDIS_ENABLE")) {
		g_conf.ConfRedis.ConfSourcce = "env"
		g_conf.ConfRedis.On = true
	} else {
		return
	}

	g_conf.ConfRedis.RedisAddress = os.Getenv("CWX_REDIS_ADDRESS")
	g_conf.ConfRedis.RedisPort = os.Getenv("CWX_REDIS_PORT")
	g_conf.ConfRedis.RedisAuthPassword = os.Getenv("CWX_REDIS_PASSWORD")
	g_conf.ConfRedis.SleepMs = g_conf.ConfStress.CacheSleepMs

}

// 解析本地参数
func (g_conf *GlobalConf) appArgsConf() {
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
