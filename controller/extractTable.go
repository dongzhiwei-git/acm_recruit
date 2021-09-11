package controller

import (
	"acm_recruit/models"
	"acm_recruit/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryExtractTable(ctx *gin.Context) {
	var extractTable,et1 models.ExtractTable
	var ets []models.ApplyForm
	if err := ctx.BindJSON(&extractTable); err != nil {
		fmt.Printf("failed bind err:%v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "格式错误！",
		})
		return
	}
	et1 = extractTable
	et, err := models.QueryExtractTable(et1)
	fmt.Println("uid:", extractTable.Uid, et.Uid)
	if et.Uid == "0" || extractTable.Uid == "0"{
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "提取码错误！",
		})
		return
	}
	if extractTable.Uid == et.Uid{
		fmt.Println("==::==")
		ets, err = models.QueryAllForm()
		util.ToExcel(ets, ctx)
	}


	if err != nil {
		fmt.Printf("failed query err:%v\n", err)
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "提取码错误！",
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"msg": "提取成功！",
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
	fmt.Println("qq:", extractTable.Uid)
	err := models.CreateExtractTable(&extractTable)
	if err != nil {
		fmt.Printf("failed create err:%v\n", err)
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "创建提取码失败！",
		})
		return
	}

	fmt.Println("qq2:", extractTable.Uid)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "创建提取码成功",
	})

}
