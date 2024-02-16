package config

import (
	"flag"
	"log"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

type Redis struct {
	Address      string
	Port         string
	AuthPassword string
}

type Mysql struct {
	Address  string
	User     string
	Password string
	Dbname   string
	Port     string
	Dsn      string
	CharSet  string
}

// 全局配置
type GlobalConf struct {
	AppConfig
	WebServer
	GinWeb
	LogTask
	Stress
	Metrics
	Blog
}

// 解析本地参数
func (g_conf *GlobalConf) getCmdArg() {
	var stressSize int
	var localAddr string

	flag.IntVar(&stressSize, "p", 0, "pressure 1: low, 2: medium, 3: high; default 0 get consul")
	flag.StringVar(&localAddr, "n", "", "local ipaddress, default: 10.x.x.x")
	flag.Parse()

	// 有设置
	if len(localAddr) != 0 {
		localAddr = unit.GetlocalIP()
		g_conf.AppConfig.LocalAddr = localAddr
	}

	// 如果没有配置 modelSize, 则使用 env 配置
	if stressSize == 0 {
		log.Println("arg modelSize 0; use env conf")
		return
	}

	// setStress(stressSize, g_conf)

}
