package metrics

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/config"
	"git.services.wait/chenwx/cwxgoweb/src/unit"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 插入数据库的数量
var InsertDB = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "write_mysql_number",
	Help: "run start insert number all",
})

// 扫描redis的数量
var ScanDelRedisNum = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "redis_scan_number",
	Help: "run start scan redis count number all",
})

// 扫描redis的数量
var RedisCmdNum = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "redis_cmd_number",
	Help: "run start cmd redis count number all",
})

// func init() {
// 	prometheus.MustRegister(metrics.InsertDB)
// 	prometheus.MustRegister(metrics.ScanDelRedisNum)
// 	prometheus.MustRegister(metrics.RedisCmdNum)
// }

func init() {
	prometheus.MustRegister(InsertDB)
	prometheus.MustRegister(ScanDelRedisNum)
	prometheus.MustRegister(RedisCmdNum)
}

func HttpServerMetrics(cfg config.Metrics) {

	if !cfg.Enable {
		fmt.Println("model: no enable Metrics")
		return
	}

	fmt.Println("model: enable Metrics")
	log.Println("metrics: start")
	localAddr := unit.GetlocalIP()
	log.Printf("metrics: local addr: http://%s:%d/metrics\n", localAddr, cfg.Port)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+strconv.Itoa(cfg.Port), nil)
}
