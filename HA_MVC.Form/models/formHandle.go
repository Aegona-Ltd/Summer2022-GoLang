package models

import (
	"HA_MVC.Form/Database"
	_ "github.com/go-sql-driver/mysql"
)

func CreateForm(form *Form) (err error) {
	err = Database.DB.Create(form).Error
	if err != nil {
		return err
	}
	return nil
}
func GetALL_Form(form *[]Form) (err error) {
	err = Database.DB.Find(form).Error
	if err != nil {
		return err
	}
	return nil
}
func GetByEmail_Handle(form *[]Form, email string) (err error) {
	err = Database.DB.Where("email = ?", email).Find(form).Error
	if err != nil {
		return err
	}
	return nil

}
func GetByName_Handle(form *[]Form, name string) (err error) {
	err = Database.DB.Where("fullname = ?", name).Find(form).Error
	if err != nil {
		return err
	}
	return nil
}
func GetOrder_Handle(form *[]Form) (err error) {
	err = Database.DB.Order("fullname desc").Find(&form).Error
	if err != nil {
		return err
	}
	return nil
}
func GetByID_Handle(form *Form, id string) (err error) {
	err = Database.DB.Where("id = ?", id).First(form).Error
	if err != nil {
		return Database.DB.Error
	}
	return nil
}
func Update_Handle(form *Form, id string) (err error) {
	Database.DB.Save(form)
	return nil
}
func GetOrderByDate_Handle(form *[]Form) (err error) {
	err = Database.DB.Order("createtime desc").Find(&form).Error
	if err != nil {
		return err
	}
	return nil
}
