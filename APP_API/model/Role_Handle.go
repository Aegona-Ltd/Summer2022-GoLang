package model

import (
	"APP_API/database"

	_ "github.com/go-sql-driver/mysql"
)

// getall
func GetAllRole_Handle(role *[]RoleModel) (err error) {
	if err = database.DB.Find(role).Error; err != nil {
		return err
	}
	return nil
}

// create
func CreateRoleHandle(role *RoleModel) (err error) {
	err = database.DB.Create(role).Error
	if err != nil {
		return err
	}
	return nil
}

// getbyid
func GetroleByIdHandle(role *RoleModel, id string) (err error) {
	err = database.DB.Where("id = ?", id).First(role).Error
	if err != nil {
		return database.DB.Error
	}
	return nil
}

// update
func UpdateroleHandle(role *RoleModel, id string) (err error) {
	// fmt.Println(form)
	database.DB.Save(role)
	return nil
}

// delete
func DeleteroleHandle(role *RoleModel, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(role)
	return nil
}
