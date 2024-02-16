package logtask

import (
	"fmt"
	"log"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/config"
	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// 日志输出任务
func LogTask(cfg config.LogTask) {
	if !cfg.Enable {
		fmt.Println("model: no enable LogTask")
		return
	}
	fmt.Println("model: enable LogTask")

	log.Println("LogTask: --- logTask start")
	for {
		rnum := unit.RandNum()
		log.Println("logTask: ", rnum)
		time.Sleep(time.Second * 20)
	}
}
