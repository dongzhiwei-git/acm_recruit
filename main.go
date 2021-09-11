package main

import (
	"acm_recruit/dao"
	"acm_recruit/models"
	"acm_recruit/routers"
	"fmt"
)

func main() {

	// connect database
	err := dao.InitMySQL()
	if err != nil {
		fmt.Printf("failed to connecte mysql: %v", err)
		return
	}

	// bind model
	// create table applyFrom
	err = dao.GetDBInstance().DB.AutoMigrate(models.ApplyForm{})
	if err != nil {
		return
	}
	err = dao.GetDBInstance().DB.AutoMigrate(models.ExtractTable{})
	if err != nil {
		return
	}
	err = dao.GetDBInstance().DB.AutoMigrate(models.Admin{})
	if err != nil {
		return
	}

	// setup router
	r := routers.SetupRouter()

	// setup listen
	err = r.Run(":8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return
	}

}
