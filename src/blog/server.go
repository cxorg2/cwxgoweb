package blog

import (
	"errors"
	"log"
	"net/http"
	net_url "net/url"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 中间件
func dbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 将 db 对象存储到请求上下文中
		c.Set("db", db)
		c.Next()
	}
}

// 校验url 并返回正确格式
func verifyPathReferer(url string) (string, error) {

	if len(url) <= 0 {
		return "", errors.New("verfiy Referer Path is null")
	}

	u, err := net_url.Parse(url)
	if err != nil {
		return "", errors.New("verfiy Referer Path Parse error")
	}

	escape_path := u.EscapedPath()

	pattern := `^/p/([a-zA-Z0-9-]+)\.html$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(escape_path)

	if len(matches) < 2 {
		return "", errors.New("verfiy Referer Path Re match error")
	}

	return escape_path, nil
}

func getPv(c *gin.Context) {
	// 从请求上下文中获取 db 对象
	dbObj := c.MustGet("db").(*docDB)

	Referer := c.GetHeader("Referer")
	log.Println(Referer)
	RefererPath, err := verifyPathReferer(Referer)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusOK, "0")
		return
	}

	log.Println(RefererPath)
	num := dbObj.GetDocPv(RefererPath)
	c.String(http.StatusOK, strconv.Itoa(num))
}

func Server(cfg BlogConf) {

	if !cfg.Enable {
		log.Println("model: no enable Blog")
		return
	}

	log.Println("model: enable Blog Api")
	log.Println("Blog: --- start")

	// db := &docDB{Dsn: cfg.Mysql.Dsn}
	db := cfg.Mysql.MysqlConf.GetSession()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 注册中间件
	r.Use(dbMiddleware(db))

	r.GET("/api/pv", getPv)

	addr := ":" + cfg.Port
	log.Printf("gin web listen server %s\n", addr)
	err := r.Run(addr)
	if err != nil {
		log.Println("listen gin web server error ", err)
	}
}
