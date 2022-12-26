package routes

import (
	"HA_MVC.Form/controller"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	{
		r.POST("form", controller.PostForm)
	}
	return r
}
