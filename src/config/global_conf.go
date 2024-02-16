package config

import (
	"log"

	"github.com/joho/godotenv"
)

var G_CONF GlobalConf

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	G_CONF.AppConfig.getConf()
	G_CONF.Stress.getConf()
	G_CONF.Blog.getConf()

	G_CONF.Metrics.getConf()
	G_CONF.WebServer.getConf()
	G_CONF.GinWeb.getConf()

	G_CONF.LogTask.getConf()

	// 获取命令行参数, 高优先级覆盖环境变量的配置
	G_CONF.getCmdArg()
}
