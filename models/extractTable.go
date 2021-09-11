package models

import (
	"acm_recruit/dao"
	"github.com/jinzhu/gorm"
)

type ExtractTable struct {
	gorm.Model
	Uid string `gorm:"type:varchar(20)" json:"password"`
}

func QueryExtractTable(et ExtractTable) (et1 ExtractTable, err error) {
	if err = dao.DB.Where("uid=?", et.Uid).First(&et1).Error; err != nil {

		return et1,err
	}
	return et1,err

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
