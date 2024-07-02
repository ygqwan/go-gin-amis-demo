package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TryGetPage(c *gin.Context) (page, pageSize int) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("perPage", "20")
	page, _ = strconv.Atoi(pageStr)
	pageSize, _ = strconv.Atoi(pageSizeStr)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return page, pageSize
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "",
		"data":   data,
	})
}
func ResponseSuccessMsg(c *gin.Context, msg string, args ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    fmt.Sprintf(msg, args...),
	})
}
func ResponseFailErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    err.Error(),
	})
}
func ResponseFailMsg(c *gin.Context, msg string, args ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    fmt.Sprintf(msg, args...),
	})
}
