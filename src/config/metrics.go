package config

import (
	"fmt"
	"os"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

type Metrics struct {
	Enable bool
	Port   int
}

func (C *Metrics) getConf() {

	if unit.IsTrue(os.Getenv("CWX_METRICS_ENABLE")) {
		C.Enable = true
	} else {
		return
	}

	port := os.Getenv("CWX_METRICS_PORT")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("get Metrics port error")
		os.Exit(1)
	}

	C.Port = portInt

}
