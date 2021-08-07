package models

import (
	"ACMZX/dao"
)

type ExtractTable struct {
	ID  uint `gorm:"AUTO_INCREMENT"`
	Uid uint `gorm:"type:in" form:"uid"`
}

func QueryExtractTable(et ExtractTable) (err error) {

	if err = dao.DB.Where("uid=?", et.Uid).First(&et).Error; err != nil {
		return err
	}
	return err

}
func CreateExtractTable(et *ExtractTable)(err error){
	err = dao.DB.Create(et).Error
	return err
}
