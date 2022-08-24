package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main1() {

	//使用默认中间件(logger 和 recovery 中间件)创建 gin 路由
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	//router.GET("/users", func(c *gin.Context) {})
	//router.GET("/user/:id", func(c *gin.Context) {})
	//router.GET("/user/:id/mmm/*action", func(c *gin.Context) {})

	router.GET("/hello/:name/*home", getting)
	//router.GET("/hello/lisi/bj", func(c *gin.Context) {})
	//router.GET("/hello/li-si/bj", func(c *gin.Context) {})

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。.
	router.Run()
	//router.Run(":3000") hardcode 端口号

}

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	//f:=func(ctx *gin.Context) {
	//	ctx.JSON(200,"halo: "+ctx.Request.URL.Path)
	//}
	engine.GET("/universe/api/v1/im/unilog/:app/*a", getting)
	//engine.GET("/universe/api/v1/im/unilog/im", f)

	//engine.POST("/universe/api/v1/im/unilog/:app/*a", f)

	engine.Run(":8989")

}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "Hello ")
}
