package main

import (
	"log"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/api"
	"git.services.wait/chenwx/cwxgoweb/src/config"
	"git.services.wait/chenwx/cwxgoweb/src/generatedata"
	"git.services.wait/chenwx/cwxgoweb/src/ginweb"
	"git.services.wait/chenwx/cwxgoweb/src/logtask"
	"git.services.wait/chenwx/cwxgoweb/src/metrics"
)

func main() {
	g_conf := config.GetGlobalConf()

	if g_conf.ConfMetrics.On {
		go metrics.HttpServerMetrics(g_conf.ConfMetrics)
	}

	if g_conf.ConfWebServer.On {
		go api.Webserver(g_conf.ConfWebServer)
	}

	if g_conf.ConfGinWeb.On {
		go ginweb.Server1(g_conf.ConfGinWeb)
	}

	if g_conf.ConfLogTask.On {
		go logtask.LogTask()
	}

	if g_conf.ConfMysql.On {
		go generatedata.MysqlTask(g_conf.ConfMysql)
	}

	if g_conf.ConfRedis.On {
		go generatedata.RedisTask(g_conf.ConfRedis)
	}

	for {
		log.Println("------ sleep 30s tmp log out")
		time.Sleep(time.Second * 30)
	}

}
