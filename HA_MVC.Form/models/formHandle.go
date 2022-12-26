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
