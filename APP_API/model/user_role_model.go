package model

import (
	"time"
)

type UserModel struct {
	UserId        int       `json:"userid" gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:userid;"`
	FullName      string    `json:"fullname" gorm:"column:fullname;"`
	CompanyName   string    `json:"companyname" gorm:"column:companyname;"`
	BusinessPhone string    `json:"businessphone" gorm:"column:businessphone;"`
	Email         string    `json:"email" gorm:"column:email;"`
	Message       string    `json:"message" gorm:"column:message;"`
	CreateTime    time.Time `json:"createtime" gorm:"column:createtime;"`
	UpdateTime    time.Time `json:"updatetime" gorm:"column:updatetime;"`
}

func (f *UserModel) TableName() string {
	return "user"
}

type RoleModel struct {
	RoleId   int    `json:"roleid" gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:roleid;"`
	RoleName string `json:"rolename"`
}

func (f *RoleModel) TableName() string {
	return "role"
}

type Role_User struct {
	Id     int `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;"`
	UserId int `json:"userid" gorm:"ForeignKey:userid;column:userid;"`
	RoleId int `json:"roleid" gorm:"ForeignKey:roleid;column:roleid;"`
}

func (f *Role_User) TableName() string {
	return "role_user"
}
