package models
type ApplyForm struct {
	Group      string `json:"group"`
	Major      string `json:"major"`
	Classes    string `json:"classes"`
	StudentNum string `json:"studentNum"`
	Name       string `json:"name"`
	PhoneNum   string `json:"phoneNum"`
	QQNum      string `json:"qqNum"`
}
