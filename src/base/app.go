package base

import (
	"flag"
	"os"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// 本地网络信息
type BaseConfig struct {
	LocalAddr string
}

func (C *BaseConfig) GetEnvConf() {
	node_ip := os.Getenv("CWX_NODE_IP")
	C.LocalAddr = node_ip
}

func (C *BaseConfig) GetCmdArgs() {

	var localAddr string
	flag.StringVar(&localAddr, "n", "", "local ipaddress, default: 10.x.x.x")

	if len(localAddr) != 0 {
		localAddr = unit.GetlocalIP()
		C.LocalAddr = localAddr
	}

}
