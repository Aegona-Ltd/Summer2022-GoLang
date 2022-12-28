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
	err = Database.DB.Where("full_name = ?", name).Find(form).Error
	if err != nil {
		return err
	}
	return nil
}
func GetOrder_Handle(form *[]Form) (err error) {
	err = Database.DB.Order("full_name desc").Find(&form).Error
	if err != nil {
		return err
	}
	return nil
}
