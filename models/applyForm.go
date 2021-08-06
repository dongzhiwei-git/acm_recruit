package models

import (
	"ACMZX/dao"
)

type ApplyForm struct {
	Group      string `json:"group"`
	Major      string `json:"major"`
	Classes    string `json:"classes"`
	StudentNum string `json:"studentNum"`
	Name       string `json:"name"`
	PhoneNum   string `json:"phoneNum"`
	QQNum      string `json:"qqNum"`
}

func CreateForm(applyForm *ApplyForm) (err error) {
	err = dao.DB.Create(&applyForm).Error
	return err

}
