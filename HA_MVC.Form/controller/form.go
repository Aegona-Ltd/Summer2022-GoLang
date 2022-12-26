package controller

import (
	"fmt"
	"net/http"

	"HA_MVC.Form/models"
	"github.com/gin-gonic/gin"
)

func PostForm(c *gin.Context) {
	var form models.Form
	c.ShouldBindJSON(&form)
	err := models.CreateForm(&form)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, form)
	}
}
