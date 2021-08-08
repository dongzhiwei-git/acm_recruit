package util

import (
	"acm_recruit/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func 	ToExcel(form []models.ApplyForm, ctx *gin.Context){
	categories := map[string]string{
		"A1": "序号",
		"B1": "组别",
		"C1": "专业",
		"D1": "学院",
		"E1": "班级",
		"F1": "学号",
		"G1": "姓名",
		"H1": "性别",
		"I1": "电话号码",
		"J1": "QQ",
		"K1": "自我介绍",

	}

	values := map[int]map[string]string{}




    for  i, v := range form{
    	fmt.Println(i+1,v)
    	values[i+1]=map[string]string{
    		"A"+strconv.Itoa(i+2):strconv.FormatUint(uint64(v.ID), 10),
    		"B"+strconv.Itoa(i+2):v.Group,
    		"C"+strconv.Itoa(i+2):v.Major,
    		"D"+strconv.Itoa(i+2):v.Introductions,
    		"E"+strconv.Itoa(i+2):v.Classes,
    		"F"+strconv.Itoa(i+2):v.StudentNum,
    		"G"+strconv.Itoa(i+2):v.Name,
    		"H"+strconv.Itoa(i+2):v.Sex,
    		"I"+strconv.Itoa(i+2):v.PhoneNum,
    		"J"+strconv.Itoa(i+2):v.QQNum,
    		"K"+strconv.Itoa(i+2):v.Introductions,

		}

	}
	f := excelize.NewFile()
	for k, v := range categories {
		err := f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return 
		}
	}
	for _, v := range values {
		for k1, v1 := range v{
			err := f.SetCellValue("Sheet1", k1, v1)
			if err != nil {
				return 
			}
		}
		
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	ctx.Header("Content-Transfer-Encoding", "binary")
	_ = f.Write(ctx.Writer)
	//err := f.SaveAs("testoo1.xlsx")
	//if err != nil {
	//	panic(err)
	//}


}
