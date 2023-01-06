package model

import (
	"time"
)

type UserModel struct {
	// gorm.Model
	UserId        int         `json:"userid" gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:userid;"`
	FullName      string      `json:"fullname" gorm:"column:fullname;"`
	CompanyName   string      `json:"companyname" gorm:"column:companyname;"`
	BusinessPhone string      `json:"businessphone" gorm:"column:businessphone;"`
	Email         string      `json:"email" gorm:"column:email;"`
	Password      string      `json:"password" gorm:"column:password;"`
	Message       string      `json:"message" gorm:"column:message;"`
	CreateTime    time.Time   `json:"createtime" gorm:"column:createtime;"`
	UpdateTime    time.Time   `json:"updatetime" gorm:"column:updatetime;"`
	Role_User     []Role_User `gorm:"foreignKey:UserID;references:userid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (f *UserModel) TableName() string {
	return "user"
}

type Role_User struct {
	Id     int    `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;"`
	UserID int    `gorm:"column:userid"`
	Name   string `json:"name"`
}

func (f *Role_User) TableName() string {
	return "role_user"
}
