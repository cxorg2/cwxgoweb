package config

import (
	"flag"
	"log"

	"git.services.wait/chenwx/cwxgoweb/src/api"
	"git.services.wait/chenwx/cwxgoweb/src/base"
	"git.services.wait/chenwx/cwxgoweb/src/blog"
	"git.services.wait/chenwx/cwxgoweb/src/generatedata"
	"git.services.wait/chenwx/cwxgoweb/src/ginweb"
	"git.services.wait/chenwx/cwxgoweb/src/logtask"
	"git.services.wait/chenwx/cwxgoweb/src/metrics"
	"git.services.wait/chenwx/cwxgoweb/src/stress"
	"github.com/joho/godotenv"
)

// 全局配置
type GlobalConf struct {
	base.BaseConfig
	api.WebServerConf
	ginweb.GinWeb
	logtask.LogTaskConf
	metrics.MetricsConf
	generatedata.GenerateConf
	blog.BlogConf
	stress.StressConf
}

var G_CONF GlobalConf

func init() {

	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading .env file")
		log.Println("no found .env file")
	}

	G_CONF.BaseConfig.GetEnvConf()
	G_CONF.GenerateConf.GetEnvConf()
	G_CONF.BlogConf.GetEnvConf()

	G_CONF.MetricsConf.GetEnvConf()
	G_CONF.WebServerConf.GetEnvConf()
	G_CONF.GinWeb.GetEnvConf()

	G_CONF.LogTaskConf.GetEnvConf()
	G_CONF.StressConf.GetEnvConf()

	// 获取命令行参数, 高优先级覆盖环境变量的配置
	G_CONF.BaseConfig.GetCmdArgs()
	G_CONF.GenerateConf.GetCmdArgs()
	G_CONF.StressConf.GetCmdArgs()
	flag.Parse()
}
