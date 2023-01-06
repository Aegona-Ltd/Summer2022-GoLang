package controller

import (
	"APP_API/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Start Skill User
// create
func Create_User(c *gin.Context) {
	log.Print("Start Create_User")
	var user model.UserModel
	c.ShouldBindJSON(&user)
	var time = time.Now()
	user.CreateTime = time
	user.UpdateTime = time
	user.Role_User = []model.Role_User{
		{Name: "customer"},
	}

	err := model.CreateUserHandle(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
	log.Print("End")
}

// update user
func Update_User(c *gin.Context) {
	log.Print("Start Update_User")
	var user model.UserModel
	id := c.Params.ByName("id")
	err := model.GetUserByIdHandle(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.ShouldBindJSON(&user)
	err = model.UpdateUserHandle(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
	log.Print("End Update_User")
}

//End Skill User

// Start Skill Admin
// get all user
func GetALL_User(c *gin.Context) {
	log.Print("Start GetALL_User")
	var role_user []model.UserModel
	c.ShouldBindJSON(&role_user)
	err := model.GetAllUser_Handle(&role_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, role_user)
	}
	log.Print("End GetALL_User")
}

// get user by userid
func GetById_User(c *gin.Context) {
	log.Print("Start GetById_User")
	var user model.UserModel
	id := c.Params.ByName("id")
	err := model.GetUserByIdHandle(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
	log.Print("End GetById_User")
}

// delete user by user id
func Delete_User(c *gin.Context) {
	log.Print("Start Delete_User")
	var user model.UserModel
	id := c.Params.ByName("id")
	err := model.DeleteUserHandle(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

	log.Print("End Delete_User")
}

// update role user
func Update_Role_User(c *gin.Context) {
	log.Print("Start Update_Role_User")
	var role_user model.Role_User
	id := c.Params.ByName("id")
	// if c.GetString("name") == "admin" {
	// 	return c.JSON(http.StatusNotFound)
	// }

	err := model.GetRole_UserByIdHandle(&role_user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.ShouldBindJSON(&role_user)
	log.Println("=========================", role_user.Name)
	if role_user.Name == "customer" || role_user.Name == "staff" {
		err = model.UpdateRole_User_Handle(&role_user, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, role_user)
		}
	} else {
		c.JSON(http.StatusUnauthorized, false)
	}
	log.Print("End Update_Role_User")
}

//End Skill Admin
