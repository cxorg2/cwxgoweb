package unit

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 获取指定 ip 前缀的本地ip地址
// TODO: 获取方式有很大的问题
func getLocalIPaddressPrefix(addHeader string) (string, error) {
	var ipString string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ipString, err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipv4 := ipnet.IP.String()
				if strings.Contains(ipv4, addHeader) {
					return ipnet.IP.String(), nil
				}
			}
		}
	}

	return "", errors.New("not found ip")
}

func GetlocalIP() (ip string) {
	ipaddr, err := getLocalIPaddressPrefix("10.")
	if err != nil {
		fmt.Println(err.Error())
		panic("get local ip err")
	}
	return ipaddr
}

// true - false 真假判断
func IsTrue(body string) bool {
	validValues := []string{"true", "Ture", "on", "ON", "YES", "yes", "y", "Y"}
	for _, value := range validValues {
		// 不区分大小写
		if strings.EqualFold(body, value) {
			return true
		}
	}
	return false
}

// 定时器, 统计1分钟执行的数量
func CountNumTicker(header string, number *int64) {
	ticker := time.NewTicker(1 * time.Minute)
	// ticker := time.NewTicker(10 * time.Second)

	var oldNumber int64
	var count int64

	for {
		<-ticker.C
		count = *number - oldNumber
		oldNumber = *number
		// log.Printf("%s: 1 min ticker insert number: +", count, "all num: ", *number)
		log.Printf("%s: 1 min ticker, number: %d all number: %d\n", header, count, *number)
	}
}
