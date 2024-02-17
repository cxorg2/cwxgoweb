package logtask

import (
	"log"
	"os"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

type LogTaskConf struct {
	Enable bool
}

func (C *LogTaskConf) GetEnvConf() {
	if unit.IsTrue(os.Getenv("CWX_LOGTASK_ON")) {
		C.Enable = true
	}
}

// 日志输出任务
func LogTask(cfg LogTaskConf) {
	if !cfg.Enable {
		log.Println("model: no enable LogTask")
		return
	}
	log.Println("model: enable LogTask")

	log.Println("LogTask: --- logTask start")
	for {
		rnum := unit.RandNum()
		log.Println("logTask: ", rnum)
		time.Sleep(time.Second * 20)
	}
}
