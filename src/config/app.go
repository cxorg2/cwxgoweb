package config

import "os"

// 本地网络信息
type AppConfig struct {
	LocalAddr string
}

func (C *AppConfig) getConf() {
	node_ip := os.Getenv("CWX_NODE_IP")
	C.LocalAddr = node_ip
}
