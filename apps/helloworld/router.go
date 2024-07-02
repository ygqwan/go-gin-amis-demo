package helloworld

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "",
		"data": gin.H{
			"title": "hello",
			"body":  "world",
		},
	})
}

func InitRouter(group *gin.RouterGroup) {
	group.GET("/hello/world", HelloWorld)
}
