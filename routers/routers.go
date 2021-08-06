package routers

import (
	"ACMZX/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	// found templates
	r.LoadHTMLGlob("templates/*")

	// found static
	r.Static("/static", "static")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "acmIndex.html", nil)
	})
	r.GET("/join", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	// submit form
	// create data
	r.POST("/join", controller.CreateForm)
	return r
}
