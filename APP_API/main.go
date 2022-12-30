package main

import (
	"APP_API/database"
	"APP_API/model"
	"APP_API/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer database.DB.Close()
	database.DB.AutoMigrate(&model.UserModel{})
	database.DB.AutoMigrate(&model.RoleModel{})
	database.DB.AutoMigrate(&model.Role_User{})
	r := routes.SetupRouter()

	//running
	r.Run(":5000")
}
