package unit

import (
	"math/rand"
	"time"
)

// 获取随机int整数
func RandNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := r.Intn(50000)
	return b
}

// 获取随机CPU使用率 float64
func RandCpu() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := r.Float64() * 100
	return b
}

// 随机 int64, 指定长度
func RandNumInt64Length(Length int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := r.Int63n(Length)
	return b
}

// 获取随机时间, 毫秒级别
func RandTimeMillisecond(ms_Leng int) time.Duration {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(ms_Leng)
	dataTime := time.Millisecond * time.Duration(randomNum)
	return dataTime
}

// 获取随机时间, 分钟级别
func RandTimeMinute(ms_Leng int) time.Duration {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(ms_Leng)
	dataTime := time.Minute * time.Duration(randomNum)
	return dataTime
}
