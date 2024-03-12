package main

import (
	"log"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/config"
	"git.services.wait/chenwx/cwxgoweb/src/ginweb"
)

func main() {
	g_conf := config.G_CONF
	log.Printf("%#v", g_conf)
	log.Printf("%#v", g_conf.GenerateConf.Redis)
	log.Printf("%#v", g_conf.GenerateConf.Mysql)

	// go metrics.HttpServerMetrics(g_conf.MetricsConf)
	// go api.Webserver(g_conf.WebServerConf)
	go ginweb.Server(g_conf.GinWeb)
	// go logtask.LogTask(g_conf.LogTaskConf)
	// go blog.Server(g_conf.BlogConf)
	// go generatedata.RunGenerateData(g_conf.GenerateConf)

	for {
		time.Sleep(time.Second * 60)
	}

}
