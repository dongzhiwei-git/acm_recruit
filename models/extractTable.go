package models

import (
	"acm_recruit/dao"
	"github.com/jinzhu/gorm"
)

type ExtractTable struct {
	gorm.Model
	Uid uint `gorm:"type:in" form:"uid"`
}

func QueryExtractTable(et ExtractTable) (err error) {

	if err = dao.DB.Where("uid=?", et.Uid).First(&et).Error; err != nil {
		return err
	}
	return err

}
//func QueryAllExtractTable(et []ExtractTable) ([]ExtractTable, error ) {
//
//	if err := dao.DB.Find(&et).Error; err != nil {
//		return nil, err
//	}
//	return et, nil
//
//}

func CreateExtractTable(et *ExtractTable)(err error){
	err = dao.DB.Create(et).Error
	return err
}
