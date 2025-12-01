package route

import (
	"net/http"
	"web_app_go/logger"

	"github.com/gin-gonic/gin"
)

// InitRouter 用来注册路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinRecovery(true), logger.GinLogger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	return r
}
