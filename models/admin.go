package models

import (
	"acm_recruit/dao"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	UserName  string `json:"user_name" gorm:"type:varchar(20)"`
	Password string `json:"password" gorm:"type:varchar(20)"`
}

func CreateAdmin(admin Admin)(err error){
	err = dao.DB.Create(&admin).Error

	return err
}

func QueryAdmin(admin Admin)(err error){
	err = dao.DB.Where("user_name=? and password=?", admin.UserName,admin.Password).First(&admin).Error
	return err
}