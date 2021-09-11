package controller

import (
	"acm_recruit/models"
	"acm_recruit/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryExtractTable(ctx *gin.Context) {
	var extractTable models.ExtractTable
	var ets []models.ApplyForm
	if err := ctx.ShouldBindJSON(&extractTable); err != nil {
		fmt.Printf("failed bind err:%v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "格式错误！",
		})
		return
	}
	//fmt.Println("ww1:",extractTable[1].Password)

	result, err := models.QueryExtractTable(extractTable)
	//err = models.QueryForm(extractTable.Password)

	if err != nil {
		fmt.Printf("failed query err:%v\n", err)
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "提取码错误！",
		})
		return
	}

	ets, err = models.QueryAllForm()
	if err != nil {
		panic(err)
	}

	if result.Password == extractTable.Password {
		util.ToExcel(ets, ctx)
		fmt.Println(ets[0].Sex)
		//ctx.JSON(http.StatusOK, gin.H{
		//	"msg": "提取成功！",
		//})
	}

	fmt.Printf("failed query err:%v\n", err)
	ctx.JSON(http.StatusForbidden, gin.H{
		"msg": "提取码错误！",
	})
}
func CreateExtractTable(ctx *gin.Context) {
	var extractTable models.ExtractTable
	if err := ctx.Bind(&extractTable); err != nil {
		fmt.Printf("failed bind err:%v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "格式错误！",
		})
		return
	}
	fmt.Println("qq:", extractTable.Password)
	err := models.CreateExtractTable(&extractTable)
	if err != nil {
		fmt.Printf("failed create err:%v\n", err)
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "创建提取码失败！",
		})
		return
	}

	fmt.Println("qq2:", extractTable.Password)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "创建提取码成功",
	})

}
