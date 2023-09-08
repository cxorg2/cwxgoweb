package logtask

import (
	"log"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

// 日志输出任务
func LogTask() {
	log.Println("task: --- logTask start")
	for {
		rnum := unit.RandNum()
		log.Println("logTask: ", rnum)
		time.Sleep(time.Second * 20)
	}
}
