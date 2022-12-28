package controller

import (
	"fmt"
	"log"
	"net/http"

	"HA_MVC.Form/models"
	"github.com/gin-gonic/gin"
)

func PostForm(c *gin.Context) {
	log.Print("Start")
	var form models.Form
	c.ShouldBindJSON(&form)
	err := models.CreateForm(&form)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}
func GetALLForm(c *gin.Context) {
	log.Print("Start")
	var form []models.Form
	// c.ShouldBindJSON(&form)
	err := models.GetALL_Form(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}
func GetBYEmail(c *gin.Context) {
	log.Print("Start")
	var form []models.Form
	email := c.Params.ByName("email")
	err := models.GetByEmail_Handle(&form, email)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}
func GetBYName(c *gin.Context) {
	log.Print("Start")
	var form []models.Form
	name := c.Params.ByName("name")
	err := models.GetByName_Handle(&form, name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}
func GetOrderEmail(c *gin.Context) {
	log.Print("Start")
	var form []models.Form
	err := models.GetOrder_Handle(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End")
}
