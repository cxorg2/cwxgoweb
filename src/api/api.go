package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"git.services.wait/chenwx/cwxgoweb/src/unit"
)

const filePath string = "/data/abc.txt"

func readfile(w http.ResponseWriter, r *http.Request) {

	bytesData, err := os.ReadFile(filePath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write(bytesData)

}

// func writefile() {

// }

// 获取请求的源ip
func getIP(w http.ResponseWriter, r *http.Request) {
	addr := r.RemoteAddr
	log.Println("rev mess getIP", addr)
	w.WriteHeader(200)
	w.Write([]byte(addr))
}

// 获取当前节点的hostname
func getHostName(w http.ResponseWriter, r *http.Request) {
	hostName := os.Getenv("HOSTNAME")
	log.Println("rev mess getIP", hostName)
	w.WriteHeader(200)
	w.Write([]byte(hostName))
}

// 延迟1s给响应
func sleep1s(w http.ResponseWriter, r *http.Request) {
	log.Println("reve mess, sleep 1s")
	time.Sleep(time.Second * 1)
	w.WriteHeader(200)
	w.Write([]byte("sleep 1s"))
}

// 输出一条日志
func oneLog(w http.ResponseWriter, r *http.Request) {
	n := unit.RandNum()
	log.Println("log +1: ", n)
	w.WriteHeader(200)
	w.Write([]byte("to log success"))
}
