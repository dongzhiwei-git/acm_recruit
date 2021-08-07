package controller

import (
	"acm_recruit/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminLogin(ctx *gin.Context) {
	var admin models.Admin
	err := ctx.BindJSON(&admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名或密码错误！不用灰心，再试一次吧！",
		})
		return
	}
	err = models.QueryAdmin(admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名或密码错误！不用灰心，再试一次吧！",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功！",
	})

}
func AdminRegister(ctx *gin.Context) {
	var admin models.Admin
	err := ctx.BindJSON(&admin)
	if err != nil {
		fmt.Printf("err1:%v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败哦！在试一次吧",
		})
		return
	}
	err = models.CreateAdmin(admin)

	if err != nil {
		fmt.Printf("err2:%v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "注册失败哦！在试一次吧",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功！",
	})

}

