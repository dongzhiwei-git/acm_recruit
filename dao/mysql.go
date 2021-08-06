package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
var(
	DB *gorm.DB
)

//	connect database
func InitMySQL()(err error){
	conn := "root:123@tcp(127.0.0.1:3306)/Blog?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", conn)
	if err != nil{
		fmt.Printf("failed to connecte mysql: %v", err)
		return
	}
	return DB.DB().Ping()


}
//	close database
func Close(){
	DB.Close()
}


