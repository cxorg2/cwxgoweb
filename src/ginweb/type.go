package ginweb

import (
	"log"
	"os"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// GinWeb 模块
type GinWeb struct {
	Enable bool
	Port   int
}

func (C *GinWeb) GetEnvConf() {
	var err error
	if unit.IsTrue(os.Getenv("CWX_GINWEB_ENABLE")) {
		C.Enable = true
	} else {
		return
	}

	port_str := os.Getenv("CWX_GINWEB_PORT")
	C.Port, err = strconv.Atoi(port_str)
	if err != nil {
		C.Port = 19004
		log.Println("no ginweb port env")
	}

}
