package main

import (
	"ACMZX/dao"
	"ACMZX/models"
	"ACMZX/routers"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// connect database
	err := dao.InitMySQL()
	if err != nil {
		fmt.Printf("failed to connecte mysql: %v", err)
		return
	}
	defer dao.Close()
	// bind model
	// create table applyFrom
	dao.DB.AutoMigrate(models.ApplyForm{})
	dao.DB.AutoMigrate(models.ExtractTable{})

	// setup router
	r := routers.SetupRouter()

    // setup listen
	err = r.Run(":8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return
	}

}
