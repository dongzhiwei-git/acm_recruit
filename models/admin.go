package models

import (
	"acm_recruit/dao"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"type:varchar(20)"`
	Password string `json:"password" gorm:"type:varchar(20)"`
}

func CreateAdmin(ctx *gin.Context, admin Admin) (err error) {

	// NOTE: 尽量用这种
	db, err := dao.GetContextDB(ctx)
	if err != nil {
		return
	}

	err = db.Create(&admin).Error

	return err
}

func QueryAdmin(admin Admin) (err error) {
	err = dao.GetDBInstance().DB.Where("user_name=? and password=?", admin.UserName, admin.Password).First(&admin).Error
	return err
}
