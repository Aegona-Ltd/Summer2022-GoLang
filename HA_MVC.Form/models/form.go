package models

import "time"

type Form struct {
	Id            int        `json:"id" gorm:"column:id;"`
	FullName      string     `json:"fullname" gorm:"column:fullname;"`
	CompanyName   string     `json:"companyname" gorm:"column:companyname;"`
	BusinessPhone string     `json:"businessphone" gorm:"column:businessphone;"`
	Email         string     `json:"email" gorm:"column:email;"`
	Message       string     `json:"message" gorm:"column:message;"`
	CreateTime    *time.Time `json:"createtime" gorm:"column:createtime;"`
	UpdateTime    *time.Time `json:"updatetime" gorm:"column:updatetime;"`
}

func (f *Form) TableName() string {
	return "form"
}
