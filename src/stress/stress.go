package stress

import (
	"crypto/sha256"
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

type StressConf struct {
	Enable   bool
	Level    int
	work_num int // 控制协程数量
}

func (C *StressConf) GetEnvConf() {
	if unit.IsTrue(os.Getenv("CWX_STRESS_ENABLE")) {
		C.Enable = true
	} else {
		return
	}

	level := os.Getenv("CWX_STRESS_LEVEL")
	if level != "" {
		C.Level, _ = strconv.Atoi(level)
	}

}

func (C *StressConf) GetCmdArgs() {
	var work_num int
	flag.IntVar(&work_num, "w", 30, "stress an int work Goroutine number")
	if work_num != 30 {
		C.work_num = work_num
	}
}

// 一轮CPU计算
func stressCPU(long_str string) {
	for i := 0; i < 1000; i++ {
		random_n := rand.Int()
		work_txt_string := long_str + strconv.Itoa(random_n)

		s_ob := sha256.New()
		s_ob.Write([]byte(work_txt_string))
		_ = s_ob.Sum(nil)
	}
}

// 单个协程
func work_task() {
	var str_txt string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 复制100份
	long_str := strings.Repeat(str_txt, 100)

	var count int = 0

	for {
		stressCPU(long_str)

		time.Sleep(time.Millisecond * 500)
		count++
		if count >= 5000 {
			break
		}
	}
}

// 核心压力任务
func RunStress(cfg StressConf) {

	if !cfg.Enable {
		log.Println("model: no enable Stress")
		return
	}

	var current_work_num int

	for {
		// 返回当前的协程数量
		current_work_num = runtime.NumGoroutine()

		if current_work_num <= cfg.work_num {
			go work_task()
		}
		time.Sleep(time.Second * 1)
	}

}
