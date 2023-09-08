package api

import (
	"log"
	"net/http"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/config"
)

func Webserver(conf config.ConfWebServer) {
	log.Println("task: --- webServer start")

	addr := ":" + strconv.Itoa(conf.Port)

	http.HandleFunc("/", rootRoute)

	log.Printf("WebServer addr: http://127.0.0.1%s/\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("服务器开启错误: ", err)
	}
}
