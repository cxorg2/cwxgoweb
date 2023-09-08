package ginweb

import (
	"log"
	"net/http"
	"strconv"

	"git.services.wait/chenwx/cwxgoweb/src/config"
	"github.com/gin-gonic/gin"
)

func Server1(conf config.ConfGinWeb) {
	r := gin.Default() // 创建实例, 这个实例即 WSGI 应用程序

	// 设置受信任的代理
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 声明路由

	// 无参数路由
	r.GET("/", root)
	r.GET("/ping", ping)

	// 路径解析
	r.GET("/user/:name", getName)

	// 参数解析
	// http://127.0.0.1:8080/users?name=chenwx
	r.GET("/users", QueryName)

	// 获取POST参数
	// curl http://localhost:8080/form  -X POST -d 'username=chenwx&password=1234'
	r.POST("/form", postArgForm)

	// Query和POST混合参数
	r.POST("/posts", postGet)
	// curl "http://localhost:8080/posts?id=9876&page=7"  -X POST -d 'username=chenwx&password=1234'

	// Map参数(字典参数)
	// curl -g "http://localhost:8080/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	r.POST("/post", postMap)

	// 重定向(Redirect)
	r.GET("/redirect", func(c *gin.Context) {
		// 301
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	// 内部代理重定向
	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.Run()

	addr := ":" + strconv.Itoa(conf.Port)
	log.Printf("gin web listen server %s\n", addr)
	err := r.Run(addr)
	if err != nil {
		log.Println("listen gin web server error ", err)
	}
}
