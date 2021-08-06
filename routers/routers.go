package routers

import (
	"ACMZX/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine{
	var applyForm models.ApplyForm
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
	r.POST("/join", func(ctx *gin.Context) {
		if err := ctx.BindJSON(&applyForm); err != nil {
			fmt.Printf("Bind failed: %v\n", err)
			return
		}
		fmt.Println(applyForm)
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	return r
}
