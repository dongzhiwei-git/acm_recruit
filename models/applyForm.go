package models

import (
	"acm_recruit/dao"
	"github.com/jinzhu/gorm"
	//"strconv"
)

type ApplyForm struct {
	gorm.Model
	Group         string `json:"group" gorm:"type:varchar(20)"`
	Name          string `json:"name" gorm:"type:varchar(20)"`
	Sex           string `json:"sex" gorm:"type:varchar(20)"`
	Series        string `json:"series" gorm:"type:varchar(20)"`
	Classes       string `json:"classes" gorm:"type:varchar(20)"`
	StudentNum    string `json:"studentNum" gorm:"type:varchar(20)"`
	QQNum         string `json:"qqNum" gorm:"type:varchar(20)"`
	PhoneNum      string `json:"phoneNum" gorm:"type:varchar(20)"`
	Introductions string `json:"introductions"`
}

func CreateForm(applyForm *ApplyForm) (err error) {
	//if sex, _ := strconv.Atoi(applyForm.Sex);sex==1{
	//	applyForm.Sex = "男"
	//}else {
	//	applyForm.Sex = "女"
	//}
	err = dao.DB.Create(&applyForm).Error
	return err

}
func QueryAllForm() (af []ApplyForm, err error) {
	err = dao.DB.Find(&af).Error
	return af, err
}
func QueryForm(uid uint) (err error) {
	err = dao.DB.Where("uid=?", uid).First(&ExtractTable{}).Error
	return err
}
