package ginweb

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Server(cfg GinWeb) {

	if !cfg.Enable {
		log.Println("model: no enable GinWeb")
		return
	}
	log.Println("model: enable GinWeb")

	// 创建实例, 这个实例即 WSGI 应用程序
	r := gin.Default()

	// 设置受信任的代理
	r.SetTrustedProxies([]string{"127.0.0.1"})

	set_route(r)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.Run()

	addr := ":" + strconv.Itoa(cfg.Port)
	log.Printf("GinWeb: listen server %s\n", addr)
	err := r.Run(addr)
	if err != nil {
		log.Println("GinWeb: listen gin web server error ", err)
	}
}
