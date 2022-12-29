package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"HA_MVC.Form/models"
	"github.com/gin-gonic/gin"
)

func PostForm(c *gin.Context) {
	log.Print("Start PostForm")
	var form models.Form
	c.ShouldBindJSON(&form)
	var time = time.Now()
	form.CreateTime = &time
	form.UpdateTime = &time
	err := models.CreateForm(&form)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End PostForm")
}
func GetALLForm(c *gin.Context) {
	log.Print("Start GetALLForm")
	var form []models.Form
	// c.ShouldBindJSON(&form)
	err := models.GetALL_Form(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End GetALLForm")
}
func GetBYEmail(c *gin.Context) {
	log.Print("Start GetBYEmail")
	var form []models.Form
	email := c.Params.ByName("email")
	err := models.GetByEmail_Handle(&form, email)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End GetBYEmail")
}
func GetBYName(c *gin.Context) {
	log.Print("Start GetBYName")
	var form []models.Form
	name := c.Params.ByName("name")
	err := models.GetByName_Handle(&form, name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End GetBYName")
}
func GetOrderEmail(c *gin.Context) {
	log.Print("Start GetOrderEmail")
	var form []models.Form
	err := models.GetOrder_Handle(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End GetOrderEmail")
}
func PutData(c *gin.Context) {
	log.Print("Start PutData")
	var form models.Form
	id := c.Params.ByName("id")
	err := models.GetByID_Handle(&form, id)
	if err != nil {
		c.JSON(http.StatusNotFound, form)
	}
	c.ShouldBindJSON(&form)
	err = models.Update_Handle(&form, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End PutData")
}
func GetOrderCreateTime(c *gin.Context) {
	log.Print("Start GetOrderCreateTime")
	var form []models.Form
	err := models.GetOrderByDate_Handle(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
	log.Print("End GetOrderCreateTime")
}
