package config

import (
	"os"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// LogTask 模块
type LogTask struct {
	Enable bool
}

func (C *LogTask) getConf() {
	if unit.IsTrue(os.Getenv("CWX_LOGTASK_ON")) {
		C.Enable = true
	}
}
