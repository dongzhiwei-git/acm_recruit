package controller

import (
	"ACMZX/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 	create an data
func CreateForm(ctx *gin.Context) {
	// gets the data from the request
	var applyForm models.ApplyForm
	if err := ctx.BindJSON(&applyForm); err != nil {
		fmt.Printf("Bind failed: %v\n", err)
		return
	}
	// logged data
    models.CreateForm(&applyForm)
	fmt.Println(applyForm)
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "恭喜您！报名成功！",
	})
}