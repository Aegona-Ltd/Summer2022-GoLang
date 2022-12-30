package controller

import (
	"APP_API/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getall
func GetALL_Role(c *gin.Context) {
	log.Print("Start GetALL_Role")
	var role []model.RoleModel
	c.ShouldBindJSON(&role)
	err := model.GetAllRole_Handle(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, role)
	}
	log.Print("End GetALL_Role")
}

// create
func Create_Role(c *gin.Context) {
	log.Print("Start Create_Role")
	var role model.RoleModel
	c.ShouldBindJSON(&role)
	log.Print("===========================", role)
	err := model.CreateRoleHandle(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, role)
	}
	log.Print("End Create_Role")
}

// getformbyid
func GetById_Role(c *gin.Context) {
	log.Print("Start GetById_Role")
	var role model.RoleModel
	id := c.Params.ByName("id")
	// c.ShouldBindJSON(&form)
	err := model.GetroleByIdHandle(&role, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, role)
	}
	log.Print("End GetById_Role")
}

// update
func Update_Role(c *gin.Context) {
	log.Print("Start Update_Role")
	var role model.RoleModel
	id := c.Params.ByName("id")
	err := model.GetroleByIdHandle(&role, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.ShouldBindJSON(&role)
	err = model.UpdateroleHandle(&role, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, role)
	}
	log.Print("End Update_Role")
}

// delete
func Delete_Role(c *gin.Context) {
	log.Print("Start Delete_Role")
	var role model.RoleModel
	id := c.Params.ByName("id")
	// c.ShouldBindJSON(&form)
	err := model.DeleteroleHandle(&role, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

	log.Print("End Delete_Role")
}
