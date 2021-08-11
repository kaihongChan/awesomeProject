package initialize

import (
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	var Router = gin.New()
	gin.SetMode(gin.DebugMode)
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return Router
}
