package main

import (
	"ACMZX/routers"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMySQL()(err error){
	conn := "root@123:tcp(127.0.0.1:3306)/Blog?charset=utf8&parseTime=True&loc=Local"
	 DB, err := gorm.Open("mysql", conn)
	if err != nil{
		fmt.Printf("failed to connecte mysql: %v", err)
		return
	}
	return DB.DB().Ping()


}

func main() {
	err := InitMySQL()


	// setup router
	r := routers.SetupRouter()

	err = r.Run("127.0.0.1:8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return
	}

}
