package model

import (
	"APP_API/database"

	_ "github.com/go-sql-driver/mysql"
)

// get all user and role
func GetAllUser_Handle(user *[]UserModel) (err error) {
	if err = database.DB.Model(&UserModel{}).Preload("Role_User").Find(&user).Error; err != nil {
		return err
	}
	return nil
}

// create
func CreateUserHandle(user *UserModel) (err error) {
	err = database.DB.Create(&user).Error
	if err == nil {
		return err
	}
	return nil
}

// getbyid
func GetUserByIdHandle(user *UserModel, id string) (err error) {
	err = database.DB.Model(&UserModel{}).Preload("Role_User").Where("userid = ?", id).First(user).Error
	if err != nil {
		return database.DB.Error
	}
	return nil
}

// update
func UpdateUserHandle(user *UserModel, id string) (err error) {
	// fmt.Println(form)
	database.DB.Save(user)
	return nil
}

// delete
func DeleteUserHandle(user *UserModel, id string) (err error) {
	database.DB.Where("userid = ?", id).Delete(user)
	return nil
}
