package main

import (
	"fmt"
	"log"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/api"
	"git.services.wait/chenwx/cwxgoweb/src/config"
	"git.services.wait/chenwx/cwxgoweb/src/ginweb"
	"git.services.wait/chenwx/cwxgoweb/src/logtask"
	"git.services.wait/chenwx/cwxgoweb/src/metrics"
)

func main() {
	g_conf := config.G_CONF
	fmt.Println(g_conf)
	fmt.Println(g_conf.Stress.Redis.Address)

	go metrics.HttpServerMetrics(g_conf.Metrics)

	go api.Webserver(g_conf.WebServer)
	go ginweb.Server1(g_conf.GinWeb)

	go logtask.LogTask(g_conf.LogTask)

	// if g_conf.ConfMysql.On {
	// 	go generatedata.MysqlTask(g_conf.ConfMysql)
	// }

	// if g_conf.ConfRedis.On {
	// 	go generatedata.RedisTask(g_conf.ConfRedis)
	// }

	// if g_conf.ConfBlogApi.On {
	// 	go blog.Server(g_conf.ConfBlogApi, g_conf.ConfBlogMysql)
	// }

	for {
		log.Println("------ sleep 30s tmp log out")
		time.Sleep(time.Second * 30)
	}

}
