package model

import (
	"APP_API/database"
)

// getbyid
func GetRole_UserByIdHandle(role_user *Role_User, id string) (err error) {
	err = database.DB.Where("userid = ?", id).First(role_user).Error
	if err != nil {
		return database.DB.Error
	}
	return nil
}

// update
func UpdateRole_User_Handle(role_user *Role_User, id string) (err error) {
	// fmt.Println(form)
	database.DB.Save(role_user)
	return nil
}
