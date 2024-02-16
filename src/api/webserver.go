package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/config"
)

func Webserver(cfg config.WebServer) {

	if !cfg.Enable {
		fmt.Println("model: no enable WebServer")
		return
	}

	fmt.Println("model: enable WebServer")
	log.Println("WebServer: --- webServer start")

	addr := ":" + strconv.Itoa(cfg.Port)

	http.HandleFunc("/", rootRoute)

	log.Printf("WebServer: addr: http://127.0.0.1%s/\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("WebServer: 服务器开启错误: ", err)
	}
}
