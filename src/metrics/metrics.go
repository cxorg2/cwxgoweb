package metrics

import (
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

func HttpServerMetrics(conf config.ConfMetrics) {

	log.Println("task: --- metrics start")
	localAddr := unit.GetlocalIP()
	log.Printf("local Metrics addr: http://%s:%d/metrics\n", localAddr, conf.MetricsPort)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+strconv.Itoa(conf.MetricsPort), nil)
}
