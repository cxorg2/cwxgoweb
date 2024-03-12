package api

import (
	"log"
	"os"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// webServer 模块
type WebServerConf struct {
	Enable bool
	Port   int
}

// 获取webserver配置
func (C *WebServerConf) GetEnvConf() {

	if unit.IsTrue(os.Getenv("CWX_WEBSERVER_ENABLE")) {
		C.Enable = true
	} else {
		return
	}

	var err error
	port_str := os.Getenv("CWX_WEBSERVER_PORT")
	C.Port, err = strconv.Atoi(port_str)
	if err != nil {
		C.Port = 19002
		log.Println("no webserver port env")
	}

}
