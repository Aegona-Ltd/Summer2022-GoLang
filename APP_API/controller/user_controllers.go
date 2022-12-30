package controller

import (
	"APP_API/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// getall
func GetALL_User(c *gin.Context) {
	log.Print("Start GetALL_User")
	var user []model.UserModel
	c.ShouldBindJSON(&user)
	err := model.GetAllUser_Handle(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
	log.Print("End GetALL_User")
}

// create
func Create_User(c *gin.Context) {
	log.Print("Start Create_User")
	var user model.UserModel
	c.ShouldBindJSON(&user)
	var time = time.Now()
	user.CreateTime = time
	user.UpdateTime = time
	err := model.CreateUserHandle(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
	log.Print("End")
}

// getformbyid
func GetById_User(c *gin.Context) {
	log.Print("Start GetById_User")
	var user model.UserModel
	id := c.Params.ByName("id")
	// c.ShouldBindJSON(&form)
	err := model.GetUserByIdHandle(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user.UserId})
	}
	log.Print("End GetById_User")
}

// update
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
		c.JSON(http.StatusOK, gin.H{"user": user.UserId})
	}
	log.Print("End Update_User")
}

// delete
func Delete_User(c *gin.Context) {
	log.Print("Start Delete_User")
	var user model.UserModel
	id := c.Params.ByName("id")
	// c.ShouldBindJSON(&form)
	err := model.DeleteUserHandle(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

	log.Print("End Delete_User")
}
