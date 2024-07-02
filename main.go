package main

import (
	"html/template"
	"log"
	"net/http"
	"tvnames/apps/helloworld"

	"github.com/gin-gonic/gin"
)

func setRouter(engine *gin.Engine) {
	// 实现静态资源服务器
	engine.Static("/web/public", "web/public")
	engine.Static("/web/pages", "web/pages")

	// 注册后端PATH
	helloworld.InitRouter(engine.Group("/api"))

	// 所有前端PATH都返回这个静态资源
	tmpl := template.Must(template.New("").ParseFiles("web/index.tmpl"))
	engine.SetHTMLTemplate(tmpl)
	engine.Use(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", nil)
	})
}

func main() {
	var engine = gin.Default()
	setRouter(engine)
	if err := engine.Run(":8080"); err != nil {
		log.Fatalln(err)
		return
	}
}
