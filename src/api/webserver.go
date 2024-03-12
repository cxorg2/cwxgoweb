package api

import (
	"log"
	"net/http"
	"strconv"
)

func Webserver(cfg WebServerConf) {

	if !cfg.Enable {
		log.Println("model: no enable WebServer")
		return
	}

	log.Println("model: enable WebServer")
	log.Println("WebServer: --- webServer start")

	addr := ":" + strconv.Itoa(cfg.Port)

	http.HandleFunc("/", rootRoute)

	log.Printf("WebServer: addr: http://127.0.0.1%s/\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("WebServer: 服务器开启错误: ", err)
	}
}
