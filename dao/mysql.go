package dao

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

var dbManager *DataManager

type DataManager struct {
	DB *gorm.DB
}

//	connect database
func InitMySQL() (err error) {
	conn := "root:root@tcp(localhost:3306)/acmrecruit?charset=utf8&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(conn))

	if err != nil {
		return fmt.Errorf("failed to connecte mysql: %v", err)
	}
	DB = DB.Set("gorm:table_options", fmt.Sprintf("ENGINE=%s CHARSET=%s COLLATE=%s", "InnoDB", "utf8", "utf8_general_ci"))

	dbManager = &DataManager{DB: DB}

	return nil
}

func GetContextDB(c *gin.Context) (*gorm.DB, error) {
	db, ok := c.Request.Context().Value("DB").(*gorm.DB)
	if ok {
		return db, nil
	} else {
		return nil, errors.New("get db error")
	}
}

func GetDBInstance() *DataManager {
	return dbManager
}
