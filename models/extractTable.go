package models

import (
	"acm_recruit/dao"
	"github.com/jinzhu/gorm"
)

type ExtractTable struct {
	gorm.Model
	Password string `json:"password" gorm:"type:varchar(20)" form:"password"`
}

func QueryExtractTable(et ExtractTable) (result ExtractTable, err error) {

	if err = dao.DB.Where("password=?", et.Password).First(&result).Error; err != nil {
		return result, err
	}
	return result, nil

}

//func QueryAllExtractTable(et []ExtractTable) ([]ExtractTable, error ) {
//
//	if err := dao.DB.Find(&et).Error; err != nil {
//		return nil, err
//	}
//	return et, nil
//
//}

func CreateExtractTable(et *ExtractTable) (err error) {
	err = dao.DB.Create(et).Error
	return err
}
