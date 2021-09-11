package routers

import (
	"acm_recruit/controller"
	"acm_recruit/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func SetupRouter() *gin.Engine {
	//initialization page

	{
		r = gin.Default()
		// to solve the cross domain
		r.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})
	}

	r.Use(middlewares.CrossHandler, middlewares.DBMiddleware)
	// found templates
	//r.LoadHTMLGlob("templates/*")
	//
	//// found static
	//r.Static("/static", "static")

	//r.GET("/", func(ctx *gin.Context) {
	//	ctx.HTML(http.StatusOK, "acmIndex.html", nil)
	//})

	// join us
	{
		//r.GET("/join", func(ctx *gin.Context) {
		//	ctx.HTML(http.StatusOK, "login.html", nil)
		//})
		// submit form
		// create data

		r.POST("/join", controller.CreateForm)
	}

	// extract table
	{
		r.POST("/extract", controller.QueryExtractTable)
		r.POST("/admin", controller.CreateExtractTable)

	}
	// admin register and login
	{
		r.POST("/login", controller.AdminLogin)
		r.POST("/register", controller.AdminRegister)
	}

	return r
}
