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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "格式有误， 重新填入",
		})
		return
	}
	// logged data
	err := models.CreateForm(&applyForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "格式有误， 重新填入",
		})
		return
	}
	fmt.Println(applyForm)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "恭喜您！报名成功！",
	})
}
